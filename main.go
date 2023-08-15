package main

import (
	apiproblem "github.com/JSainsburyPLC/go-api-problem"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	nrw "github.com/JSainsburyPLC/go-newrelic-wrapper"
	"github.com/go-chi/chi/v5"

	"github.com/JSainsburyPLC/smartshop-api-go-template/app"
	"github.com/JSainsburyPLC/smartshop-api-go-template/config"
	//TODO Uncomment if you have a database "github.com/JSainsburyPLC/smartshop-api-go-template/db"
)

const (
	errorType = "https://github.com/JSainsburyPLC/smartshop-api-go-template/blob/develop/README_TEMPLATE.md#Error-Codes"
)

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		panic("could not load a config")
	}

	log.Enable(cfg.Logger.LogLevel)

	enableNewRelic(err, cfg)

	a := app.NewApplication(chi.NewRouter(), apiproblem.NewFactory(errorType))

	a.Init()
	a.Run(cfg.Port, cfg.Logger.LogHttpBodies)
}

func enableNewRelic(err error, cfg config.AppConfig) {
	if !cfg.NewRelic.Enabled {
		return
	}

	err = nrw.Enable(nrw.WrapperConfig{
		AppName: cfg.NewRelic.AppName,
		Licence: cfg.NewRelic.LicenceKey,
		Labels: map[string]string{
			"environment": cfg.NewRelic.LabelEnvironment,
			"account":     cfg.NewRelic.LabelAccount,
			"role":        cfg.NewRelic.LabelRole,
		},
		Enabled:                  cfg.NewRelic.Enabled,
		EnsureConnection:         false,
		DistributedTracerEnabled: true,
	})

	if err != nil {
		log.Errorf("failed to enable new relic: %s", err.Error())
	}
}
