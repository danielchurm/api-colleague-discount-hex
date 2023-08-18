package http_handlers_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/inbound/http_handlers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("healthcheck handler", func() {

	It("returns 200 and status ok", func() {
		request, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
		response := httptest.NewRecorder()
		expectedBody := `{"status":"OK","errors":[]}`

		healthcheckHandler := http_handlers.Healthcheck{}
		healthcheckHandler.ServeHTTP(response, request)

		Expect(response.Code).To(Equal(http.StatusOK))
		Expect(response.Body).To(MatchJSON(expectedBody))
	})
})
