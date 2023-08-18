package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/JSainsburyPLC/go-logrus-wrapper"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/healthcheck"
)

type HealthCheckHandler struct {
	hcs healthcheck.HealthChecker
}

func NewHealthCheck(hcs healthcheck.HealthChecker) HealthCheckHandler {
	return HealthCheckHandler{hcs}
}

func (h HealthCheckHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	health, _ := h.hcs.GetHealth()
	body, _ := json.Marshal(health)
	_, err := w.Write(body)

	if err != nil {
		log.Error("failed to marshal health to JSON")
	}
}
