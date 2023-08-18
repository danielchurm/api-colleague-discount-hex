package app_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	apiproblem "github.com/JSainsburyPLC/go-api-problem"
	"github.com/go-chi/chi/v5"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/app"
)

var (
	mockCtrl *gomock.Controller
)

var _ = Describe("App", func() {

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		defer mockCtrl.Finish()
	})

	Describe("Application Health", func() {
		Context("When the service is healthy", func() {
			It("reports no errors", func() {
				router := chi.NewRouter()

				a := app.NewApplication(router, apiproblem.NewFactory("http://error.type/"))
				a.Init()

				req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
				recorder := httptest.NewRecorder()

				router.ServeHTTP(recorder, req)

				body, _ := io.ReadAll(recorder.Body)

				Expect(recorder.Code).To(Equal(http.StatusOK))
				Expect(string(body)).To(Equal(`{"status":"OK","errors":[]}`))
			})
		})
	})

	Describe("Resources Not Found", func() {
		Context("When A resource cannot be found", func() {
			It("inform the consumer of it not being found", func() {
				router := chi.NewRouter()

				a := app.NewApplication(router, apiproblem.NewFactory("http://error.type/"))
				a.Init()

				req, _ := http.NewRequest(http.MethodGet, "/will/never/exist", nil)
				recorder := httptest.NewRecorder()

				router.ServeHTTP(recorder, req)

				body, _ := io.ReadAll(recorder.Body)

				expBody := `{
					"type": "http://error.type/",
					"status": 404,
					"title" : "Not Found",
					"detail": "Unable to locate resource /will/never/exist",
					"code"  : 99000
				}`

				Expect(recorder.Code).To(Equal(http.StatusNotFound))
				Expect(string(body)).To(MatchJSON(expBody))
			})
		})
	})
})
