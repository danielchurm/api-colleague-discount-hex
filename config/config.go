package config

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	Port      string `env:"PORT" envDefault:"8080"`
	ErrorType string `env:"ERROR_TYPE_URL"`
	Logger    LoggerConfig
	NewRelic  NewRelicConfig
}

type LoggerConfig struct {
	LogLevel      string `env:"LOGS_APP_LEVEL" envDefault:"info"`
	LogHttpBodies bool   `env:"LOG_HTTP_BODIES" envDefault:"false"`
}

type NewRelicConfig struct {
	AppName          string `env:"NEW_RELIC_APP_NAME" envDefault:""`
	LicenceKey       string `env:"NEW_RELIC_LICENCE_KEY" envDefault:""`
	Enabled          bool   `env:"NEW_RELIC_ENABLED" envDefault:"false"`
	LabelEnvironment string `env:"NEW_RELIC_LABEL_ENV" envDefault:"local"`
	LabelAccount     string `env:"NEW_RELIC_LABEL_ACCOUNT" envDefault:""`
	LabelRole        string `env:"NEW_RELIC_LABEL_ROLE" envDefault:""`
}

type DatabaseConfig struct {
	User           string `env:"DB_USER" envDefault:"api_go_template"`
	Name           string `env:"DB_NAME" envDefault:"api_go_template"`
	Password       string `env:"DB_PASSWORD" envDefault:"password"`
	Host           string `env:"DB_HOST" envDefault:"localhost"`
	Port           int    `env:"DB_PORT" envDefault:"5432"`
	SSLMode        string `env:"DB_SSL_MODE" envDefault:"disable"`
	Type           string `env:"DB_TYPE" envDefault:"postgres"`
	TimeoutSeconds int    `env:"DB_TIMEOUT_SECONDS" envDefault:"10"`
}

func NewAppConfig() (AppConfig, error) {
	cfg := &AppConfig{}
	if err := env.Parse(cfg); err != nil {
		return AppConfig{}, err
	}
	return *cfg, nil
}
