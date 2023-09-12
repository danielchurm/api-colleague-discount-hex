package config

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	Port                             string `env:"PORT" envDefault:"8080"`
	ErrorType                        string `env:"ERROR_TYPE_URL"`
	IdentityOrchestratorConfig       IdentityOrchestratorConfig
	CheckoutsColleagueDiscountConfig CheckoutsColleagueDiscountConfig
	Logger                           LoggerConfig
	NewRelic                         NewRelicConfig
}

type IdentityOrchestratorConfig struct {
	Host   string `env:"IDENTITY_ORCHESTRATOR_HOST" envDefault:"smartshop-api-identity-orchestrator-mock-server.app.internal"`
	ApiKey string `env:"IDENTITY_ORCHESTRATOR_API_KEY" envDefault:"the-orchestrator-api-key"`
}

type CheckoutsColleagueDiscountConfig struct {
	Scheme string `env:"CHECKOUTS_COLLEAGUE_DISCOUNT_SCHEME" envDefault:"http"`
	Host   string `env:"CHECKOUTS_COLLEAGUE_DISCOUNT_HOST" envDefault:"sainsburys-colleague-discount-mock-server.app.internal"`
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

func NewAppConfig() (AppConfig, error) {
	cfg := &AppConfig{}
	if err := env.Parse(cfg); err != nil {
		return AppConfig{}, err
	}
	return *cfg, nil
}
