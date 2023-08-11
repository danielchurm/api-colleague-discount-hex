package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/JSainsburyPLC/smartshop-api-go-template/handlers"
	"github.com/JSainsburyPLC/smartshop-api-go-template/healthcheck"
	mock_healthcheck "github.com/JSainsburyPLC/smartshop-api-go-template/mocks/healthcheck"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	mockCtrl *gomock.Controller
)

var _ = Describe("Health check handler", func() {

	var mockHealthCheckService *mock_healthcheck.MockHealthChecker

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockHealthCheckService = mock_healthcheck.NewMockHealthChecker(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("When service is healthy", func() {
		It("returns a healthy response", func() {
			mockHealthCheckService.EXPECT().GetHealth().Times(1).Return(
				healthcheck.Health{Status: "OK", Errors: []error{}},
				nil)

			h := handlers.NewHealthCheck(mockHealthCheckService)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)

			h.ServeHTTP(resp, req)

			By("Having HTTP Status OK", func() {
				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			By("Reporting status OK and no errors", func() {
				body, _ := io.ReadAll(resp.Body)
				Expect(string(body)).To(Equal(`{"status":"OK","errors":[]}`))
			})
		})
	})
})
