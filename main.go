package main

import (
	apiproblem "github.com/JSainsburyPLC/go-api-problem"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	nrw "github.com/JSainsburyPLC/go-newrelic-wrapper"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/config"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/inbound"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/outbound"
	//TODO Uncomment if you have a database "github.com/JSainsburyPLC/smartshop-api-colleague-discount/db"
)

const (
	errorType = "https://github.com/JSainsburyPLC/smartshop-api-colleague-discount/blob/develop/README_TEMPLATE.md#Error-Codes"
)

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		panic("could not load a config")
	}

	log.Enable(cfg.Logger.LogLevel)

	enableNewRelic(err, cfg)

	// Outbound adapters
	httpClient := nrw.NewHttpClient()
	identityOrchestratorClient := outbound.NewIdentityOrchestratorClient(
		cfg.IdentityOrchestratorConfig.Host,
		cfg.IdentityOrchestratorConfig.ApiKey,
		httpClient,
	)
	checkoutsColleagueDiscountClient := outbound.NewCheckoutsColleagueDiscountClient(
		cfg.CheckoutsColleagueDiscountConfig.Scheme,
		cfg.CheckoutsColleagueDiscountConfig.Host,
		httpClient,
	)

	// Domain
	discountCard := domain.NewDiscountCard(identityOrchestratorClient, checkoutsColleagueDiscountClient)

	// Inbound adapters
	apiProblemFactory := apiproblem.NewFactory(errorType)
	server := inbound.NewServer(cfg.Port, cfg.Logger.LogHttpBodies, apiProblemFactory, discountCard)
	err = server.ListenAndServe()
	if err != nil {
		log.Errorf("server exited with error %+v", err)
	}
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
