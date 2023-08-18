package http_handlers

import (
	"net/http"

	log "github.com/JSainsburyPLC/go-logrus-wrapper"
)

type Healthcheck struct{}

func (h Healthcheck) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body := `{"status": "OK", "errors": []}`
	_, err := writer.Write([]byte(body))
	if err != nil {
		log.CtxErrorf(request.Context(), "failed to write http body for healthcheck. %s", err)
	}
}
