package inbound

import (
	"fmt"
	"net/http"
	"os"
	"time"

	apiproblem "github.com/JSainsburyPLC/go-api-problem"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"github.com/JSainsburyPLC/go-logrus-wrapper/middleware"
	nrw "github.com/JSainsburyPLC/go-newrelic-wrapper"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/inbound/http_handlers"
	accesslogger "github.com/JSainsburyPLC/smartshop-go-access-logger"
	"github.com/go-chi/chi/v5"
)

const (
	errorType = "https://github.com/JSainsburyPLC/smartshop-api-colleague-discount/blob/develop/README_TEMPLATE.md#Error-Codes"
)

type Server struct {
	Port              string
	LogHttpBodies     bool
	ApiProblemFactory apiproblem.Factory
	DiscountCard      domain.DiscountCardRetriever
}

func NewServer(port string,
	logHttpBodies bool,
	apiProblemFactory apiproblem.Factory,
	discountCard domain.DiscountCardRetriever,
) Server {
	return Server{
		Port:              port,
		LogHttpBodies:     logHttpBodies,
		ApiProblemFactory: apiProblemFactory,
		DiscountCard:      discountCard,
	}
}

func (s Server) ListenAndServe() error {
	const forDockerUseGlobalIP = "0.0.0.0"

	log.Infof("app is running on IP %s, port %s", forDockerUseGlobalIP, s.Port)

	handler := s.createRouter(s.LogHttpBodies)
	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%s", forDockerUseGlobalIP, s.Port),
		Handler:      handler,
		IdleTimeout:  65 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s Server) createRouter(logHttpBodies bool) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.AddLoggableHeadersToContext)
	router.Use(nrw.InboundTransactionMiddleware)

	healthcheckHandler := http_handlers.Healthcheck{}.ServeHTTP
	router.Get("/healthcheck", healthcheckHandler)

	apiProblemFactory := apiproblem.NewFactory(errorType)

	getCardHandler := http_handlers.NewGetCardHandler(apiProblemFactory, s.DiscountCard)
	router.Get("/discount-card", getCardHandler.ServeHTTP)

	notFounderHandler := http_handlers.NewNotFound(apiProblemFactory).ServeHTTP
	router.NotFound(notFounderHandler)

	accessLoggerHandler := accesslogger.New(os.Stdout, logHttpBodies)
	return accessLoggerHandler.Handler(router)
}
