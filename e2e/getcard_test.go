package e2e_test

import (
	"fmt"
	"io"
	"net/http"

	"github.com/churmd/smockerclient/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Card", func() {
	When("requesting a valid card", func() {
		It("returns card number and issue number", func() {
			userId := "1234"
			email := "colleague1@example.com"
			cardNumber := "1000000000000001"
			issueNumber := "15"
			status := "VERIFIED"

			mockSetupCardForUser(userId, email, cardNumber, issueNumber, status)

			url := fmt.Sprintf("%s/discount-card?user_id=%s", cfg.ApiColleagueDiscountHost, userId)

			expectedBody := `{
				"card_number": "` + cardNumber + `",
				"issue_number": "` + issueNumber + `",
				"status": "` + status + `"
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

func mockSetupCardForUser(userId, email, cardNumber, issueNumber, status string) {
	Expect(mockIdentityOrchestrator.ResetAllSessionsAndMocks()).To(Succeed())
	Expect(mockSainsColleagueDiscount.ResetAllSessionsAndMocks()).To(Succeed())

	ioReq := mock.NewRequestBuilder(http.MethodGet, "/api/v1/users/"+userId).
		AddHeader("X-Api-Key", "the-orchestrator-api-key").
		Build()
	ioRespJson := `{
		"email": "` + email + `",
		"identity_uuid": "some-1-identity-uuid",
		"id": 8,
		"nectar_card": "24871760811",
		"terms_accepted": "2015-03-27T14:03:26+00:00",
		"registration_date": "2015-03-27T14:03:26+00:00"
	}`
	ioResp := mock.NewResponseBuilder(http.StatusOK).AddBody(ioRespJson).Build()
	ioMockDef := mock.NewDefinition(ioReq, ioResp)

	Expect(mockIdentityOrchestrator.AddMock(ioMockDef)).To(Succeed())

	scdreq := mock.NewRequestBuilder(http.MethodGet, "/discount-card").
		AddQueryParam("email", email).
		Build()
	scdRespJson := `{"cardNumber": "` + cardNumber + `","discountPercentage": 15,"issueNumber": ` + issueNumber + `,"status": "` + status + `"}`
	scdResp := mock.NewResponseBuilder(http.StatusOK).AddBody(scdRespJson).Build()
	scdMockDef := mock.NewDefinition(scdreq, scdResp)

	Expect(mockSainsColleagueDiscount.AddMock(scdMockDef)).To(Succeed())
}
