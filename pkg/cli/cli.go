package cli

import (
	"context"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/kernel"
	"github.com/applike/gosoline/pkg/mon"
)

type Module interface {
	Boot(config cfg.Config, logger mon.Logger) error
	Run(ctx context.Context) error
}

type cliModule struct {
	kernel.EssentialModule
	Module
}

func newCliModule(module Module) *cliModule {
	return &cliModule{
		Module: module,
	}
}

func Run(module Module) {
	configOptions := []cfg.Option{
		cfg.WithErrorHandlers(defaultErrorHandler),
		cfg.WithConfigFile("./config.dist.yml", "yml"),
		cfg.WithConfigFileFlag("config"),
	}

	config := cfg.New()
	if err := config.Option(configOptions...); err != nil {
		defaultErrorHandler(err, "can not initialize the config")
	}

	logger, err := newCliLogger()
	if err != nil {
		defaultErrorHandler(err, "can not initialize the logger")
	}

	if err := module.Boot(config, logger); err != nil {
		defaultErrorHandler(err, "can not boot the module")
	}

	settings := &kernel.Settings{}
	config.UnmarshalKey("kernel", settings)

	k := kernel.New(config, logger, settings)
	k.Add("cli", newCliModule(module))
	k.Run()
}
