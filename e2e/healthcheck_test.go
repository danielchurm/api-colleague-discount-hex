package e2e_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health Checks", func() {
	Context("Test healthy service", func() {
		It("reports that the service is healthy", func() {
			resp, err := http.Get(cfg.ApiColleagueDiscountHost + "/healthcheck")

			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
