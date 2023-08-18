package http_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JSainsburyPLC/go-api-problem"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
)

type NotFoundHandler struct {
	factory apiproblem.Factory
}

func NewNotFound(factory apiproblem.Factory) NotFoundHandler {
	return NotFoundHandler{factory: factory}
}

func (h NotFoundHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	detail := fmt.Sprintf("Unable to locate resource %s", request.URL.Path)
	status := http.StatusNotFound

	p := h.factory.NewApiProblem(status, NotFoundError, detail)

	b, _ := json.Marshal(p)

	writer.WriteHeader(status)
	_, err := writer.Write(b)
	if err != nil {
		log.CtxErrorf(request.Context(), "failed to write http body for not found handler. %s", err)
	}
}
