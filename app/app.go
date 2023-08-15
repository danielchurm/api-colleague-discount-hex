package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	apiproblem "github.com/JSainsburyPLC/go-api-problem"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"github.com/JSainsburyPLC/go-logrus-wrapper/middleware"
	nrw "github.com/JSainsburyPLC/go-newrelic-wrapper"
	accesslogger "github.com/JSainsburyPLC/smartshop-go-access-logger"
	"github.com/go-chi/chi/v5"

	"github.com/JSainsburyPLC/smartshop-api-go-template/handlers"
	"github.com/JSainsburyPLC/smartshop-api-go-template/healthcheck"
)

type App struct {
	r              chi.Router
	problemFactory apiproblem.Factory
}

func NewApplication(r chi.Router, problemFactory apiproblem.Factory) App {
	return App{r, problemFactory}
}

func (a *App) Init() {
	a.r.Use(middleware.AddLoggableHeadersToContext)
	a.r.Use(nrw.InboundTransactionMiddleware)

	a.r.Get("/healthcheck", handlers.NewHealthCheck(healthcheck.NewHealthCheckService()).ServeHTTP)

	a.r.NotFound(handlers.NewNotFound(a.problemFactory).ServeHTTP)
}

func (a *App) Run(port string, logHttpBodies bool) {
	const forDockerUseGlobalIP = "0.0.0.0"

	log.Infof("app is running on IP %s, port %s", forDockerUseGlobalIP, port)

	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%s", forDockerUseGlobalIP, port),
		Handler:      accessLogMiddleware(a.r, logHttpBodies),
		IdleTimeout:  65 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	srv.ListenAndServe()
}

func accessLogMiddleware(routingHandler http.Handler, logHttpBodies bool) http.Handler {
	accessLoggerMiddleware := accesslogger.New(os.Stdout, logHttpBodies)
	return accessLoggerMiddleware.Handler(routingHandler)
}
