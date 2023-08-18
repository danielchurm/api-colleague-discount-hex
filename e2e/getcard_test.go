package e2e_test

import (
	"fmt"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Card", func() {
	When("requesting a valid card", func() {
		It("returns card number and issue number", func() {
			userId := "1234"
			url := fmt.Sprintf("%s/discount-card?user_id=%s", cfg.ApiColleagueDiscountHost, userId)

			expectedBody := `{
				"card_number": "1000000000000001",
				"issue_number": "15",
				"status": "VERIFIED"
			}`

			resp, err := http.Get(url)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			body, err := io.ReadAll(resp.Body)
			Expect(err).To(BeNil())
			Expect(body).To(MatchJSON(expectedBody))
		})
	})
})
