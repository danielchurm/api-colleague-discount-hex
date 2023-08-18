package http_handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	apiproblem "github.com/JSainsburyPLC/go-api-problem"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/handlers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Not Found handler", func() {

	Context("When a resource cannot be found", func() {
		It("returns a descriptive error", func() {

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/roswell/alien", nil)

			h := handlers.NewNotFound(apiproblem.NewFactory("http://error.type/"))
			h.ServeHTTP(resp, req)

			By("Having HTTP Status OK", func() {
				Expect(resp.Code).To(Equal(http.StatusNotFound))
			})

			expBody := `{
					"type": "http://error.type/",
					"status": 404,
					"title" : "Not Found",
					"detail": "Unable to locate resource /roswell/alien",
					"code"  : 99000
				}`

			By("Reporting status OK and no errors", func() {
				body, _ := io.ReadAll(resp.Body)
				Expect(string(body)).To(MatchJSON(expBody))
			})
		})
	})
})
