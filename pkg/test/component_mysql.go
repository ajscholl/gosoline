package test

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/pkg/errors"
)

type mysqlSettingsLegacy struct {
	*mockSettings
	Version string `cfg:"version"`
	DbName  string `cfg:"dbName"`
}

type mysqlComponentLegacy struct {
	mockComponentBase
	settings *mysqlSettingsLegacy
	db       *sql.DB
}

func (m *mysqlComponentLegacy) Boot(config cfg.Config, _ log.Logger, runner *dockerRunnerLegacy, settings *mockSettings, name string) {
	m.name = name
	m.runner = runner
	m.settings = &mysqlSettingsLegacy{
		mockSettings: settings,
	}
	key := fmt.Sprintf("mocks.%s", name)
	config.UnmarshalKey(key, m.settings)
}

func (m *mysqlComponentLegacy) getContainerConfig() *containerConfigLegacy {
	return &containerConfigLegacy{
		Repository: "mysql/mysql-server",
		Tmpfs:      m.settings.Tmpfs,
		Tag:        m.settings.Version,
		Env: []string{
			fmt.Sprintf("MYSQL_DATABASE=%s", m.settings.DbName),
			"MYSQL_USER=gosoline",
			"MYSQL_PASSWORD=gosoline",
			"MYSQL_ROOT_PASSWORD=gosoline",
			fmt.Sprintf("MYSQL_ROOT_HOST=%s", "%"),
		},
		Cmd: []string{"--sql_mode=NO_ENGINE_SUBSTITUTION", "--log-bin-trust-function-creators=TRUE"},
		PortBindings: portBindingLegacy{
			"3306/tcp": fmt.Sprint(m.settings.Port),
		},
		PortMappings: portMappingLegacy{
			"3306/tcp": &m.settings.Port,
		},
		HostMapping: hostMappingLegacy{
			dialPort: &m.settings.Port,
			setHost:  &m.settings.Host,
		},
		HealthCheck: func() error {
			client, err := m.provideMysqlClient()
			if err != nil {
				return err
			}

			err = client.Ping()

			if err != nil {
				return err
			}

			return nil
		},
		PrintLogs:   m.settings.Debug,
		ExpireAfter: m.settings.ExpireAfter,
	}
}

func (m *mysqlComponentLegacy) PullContainerImage() error {
	containerName := fmt.Sprintf("gosoline_test_mysql_%s", m.name)
	containerConfig := m.getContainerConfig()

	return m.runner.PullContainerImage(containerName, containerConfig)
}

func (m *mysqlComponentLegacy) Start() error {
	if len(m.settings.Tmpfs) == 0 {
		m.settings.Tmpfs["/var/lib/mysql"] = ""
	}

	containerName := fmt.Sprintf("gosoline_test_mysql_%s", m.name)
	containerConfig := m.getContainerConfig()

	return m.runner.Run(containerName, containerConfig)
}

type noopLogger struct{}

func (l noopLogger) Print(v ...interface{}) {
}

func init() {
	err := mysql.SetLogger(&noopLogger{})
	if err != nil {
		panic(err)
	}
}

func (m *mysqlComponentLegacy) provideMysqlClient() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", "gosoline", "gosoline", m.settings.Host, m.settings.Port, m.settings.DbName)

	if m.db == nil {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, errors.Wrapf(err, "can not open mysql connection %s", dsn)
		}

		if db != nil {
			m.db = db
		}
	}

	return m.db, nil
}
