package suite

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/justtrackio/gosoline/pkg/stream"
	"github.com/justtrackio/gosoline/pkg/test/env"
	"github.com/stretchr/testify/assert"
)

type (
	testCaseMatcher func(method reflect.Method) bool
	testCaseBuilder func(suite TestingSuite, method reflect.Method) (testCaseRunner, error)
	testCaseRunner  func(t *testing.T, suite TestingSuite, suiteOptions *suiteOptions, environment *env.Environment)
)

type testCaseDefinition struct {
	matcher testCaseMatcher
	builder testCaseBuilder
}

var testCaseDefinitions = map[string]testCaseDefinition{}

func Run(t *testing.T, suite TestingSuite, extraOptions ...Option) {
	suite.SetT(t)

	var err error
	var testCases map[string]testCaseRunner
	suiteOptions := suiteApplyOptions(suite, extraOptions)

	if testCases, err = suiteFindTestCases(t, suite, suiteOptions); err != nil {
		assert.FailNow(t, err.Error())

		return
	}

	if len(testCases) == 0 {
		return
	}

	for i := 0; i < suiteOptions.testCaseRepeatCount; i++ {
		if suiteOptions.envIsShared {
			runTestCaseWithSharedEnvironment(t, suite, suiteOptions, testCases)
		} else {
			runTestCaseWithIsolatedEnvironment(t, suite, suiteOptions, testCases)
		}
	}
}

func suiteFindTestCases(_ *testing.T, suite TestingSuite, options *suiteOptions) (map[string]testCaseRunner, error) {
	var err error
	testCases := make(map[string]testCaseRunner)
	methodFinder := reflect.TypeOf(suite)

	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)

		if !strings.HasPrefix(method.Name, "Test") {
			continue
		}

		for typ, def := range testCaseDefinitions {
			if !def.matcher(method) {
				continue
			}

			if options.shouldSkip(method.Name) {
				testCases[method.Name] = func(t *testing.T, _ TestingSuite, _ *suiteOptions, _ *env.Environment) {
					t.SkipNow()
				}

				continue
			}

			if testCases[method.Name], err = def.builder(suite, method); err != nil {
				return nil, fmt.Errorf("can not build test case %s of type %s: %w", method.Name, typ, err)
			}
		}
	}

	if len(testCases) == 0 {
		return nil, fmt.Errorf("no testcases found. the function name has to start with 'Test'")
	}

	return testCases, nil
}

func suiteApplyOptions(suite TestingSuite, extraOptions []Option) *suiteOptions {
	setupOptions := []Option{
		WithClockProvider(clock.NewFakeClock()),
		WithConfigMap(map[string]any{
			"cloud.aws.default.ec2.metadata.available": false,
		}),
	}
	setupOptions = append(setupOptions, suite.SetupSuite()...)
	setupOptions = append(setupOptions, extraOptions...)

	options := newSuiteOptions()

	for _, opt := range setupOptions {
		opt(options)
	}

	return options
}

func runTestCaseWithSharedEnvironment(t *testing.T, suite TestingSuite, suiteOptions *suiteOptions, testCases map[string]testCaseRunner) {
	envOptions := []env.Option{
		env.WithConfigEnvKeyReplacer(cfg.DefaultEnvKeyReplacer),
		env.WithConfigMap(map[string]any{
			"app_project": "justtrack",
			"app_family":  "gosoline",
			"app_group":   "test",
			"app_name":    "test",
		}),
	}
	envOptions = append(envOptions, suiteOptions.envOptions...)
	envOptions = append(envOptions, env.WithConfigMap(map[string]any{
		"env":              "test",
		"fixtures.enabled": true,
	}))

	environment, err := env.NewEnvironment(t, envOptions...)
	if err != nil {
		assert.FailNow(t, "failed to create test environment", err.Error())
	}

	defer func() {
		if err := environment.Stop(); err != nil {
			assert.FailNow(t, "failed to stop test environment", err.Error())
		}
	}()

	suite.SetEnv(environment)

	for _, envSetup := range suiteOptions.envSetup {
		if err := envSetup(); err != nil {
			assert.FailNow(t, "failed to execute additional environment setup", err.Error())
		}
	}

	start := time.Now()
	if err = environment.LoadFixtureSets(suiteOptions.fixtureSetFactories...); err != nil {
		assert.FailNow(t, "failed to load fixtures from factories", err.Error())
	}

	environment.Logger().WithChannel("fixtures").Debug("loaded fixtures in %s", time.Since(start))

	for name, testCase := range testCases {
		runTestCaseInSuite(suite, func() {
			t.Run(name, func(t *testing.T) {
				parentT := suite.T()
				suite.SetT(t)

				testCase(t, suite, suiteOptions, environment)

				suite.SetT(parentT)
			})
		})
	}
}

func runTestCaseWithIsolatedEnvironment(t *testing.T, suite TestingSuite, suiteOptions *suiteOptions, testCases map[string]testCaseRunner) {
	for name, testCase := range testCases {
		runTestCaseWithSharedEnvironment(t, suite, suiteOptions, map[string]testCaseRunner{
			name: testCase,
		})
	}
}

func runTestCaseInSuite(suite TestingSuite, testCase func()) {
	if setupTestAware, ok := suite.(TestingSuiteSetupTestAware); ok {
		if err := setupTestAware.SetupTest(); err != nil {
			assert.FailNow(suite.T(), "failed to setup the test", err.Error())
		}
	}

	testCase()

	if tearDownTestAware, ok := suite.(TestingSuiteTearDownTestAware); ok {
		if err := tearDownTestAware.TearDownTest(); err != nil {
			assert.FailNow(suite.T(), "failed to tear down the test", err.Error())
		}
	}

	stream.ResetInMemoryInputs()
	stream.ResetInMemoryOutputs()
	suite.Env().ResetLogs()
}
