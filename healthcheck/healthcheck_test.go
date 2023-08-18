package healthcheck_test

import (
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/healthcheck"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	mockCtrl *gomock.Controller
)

var _ = Describe("Health check service", func() {

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("When service is healthy", func() {
		It("returns a healthy response", func() {
			hcs := healthcheck.NewHealthCheckService()

			health, err := hcs.GetHealth()

			By("returning Status OK", func() {
				Expect(health.Status).To(Equal("OK"))
			})

			By("returning no errors reported", func() {
				Expect(health.Errors).To(Equal([]error{}))
			})

			By("not generating an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
