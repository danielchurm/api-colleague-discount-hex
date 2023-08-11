package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JSainsburyPLC/go-api-problem"
)

type NotFoundHandler struct {
	f apiproblem.Factory
}

func NewNotFound(f apiproblem.Factory) NotFoundHandler {
	return NotFoundHandler{f}
}

func (h NotFoundHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	detail := fmt.Sprintf("Unable to locate resource %s", req.URL.Path)
	status := http.StatusNotFound

	p := h.f.NewApiProblem(status, NotFoundError, detail)

	b, _ := json.Marshal(p)

	w.WriteHeader(status)
	w.Write(b)
}
