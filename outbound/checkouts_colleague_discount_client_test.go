package outbound_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
	mock_outbound "github.com/JSainsburyPLC/smartshop-api-colleague-discount/mocks/outbound"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/outbound"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("checkouts colleague discount client", func() {
	var (
		ctrl           *gomock.Controller
		mockHttpClient *mock_outbound.MockHttpClient

		host  = "checkouts-colleague-discount.int.stg.jspaas.uk"
		email = "user123@example.com"

		ctx context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockHttpClient = mock_outbound.NewMockHttpClient(ctrl)

		ctx = context.Background()
	})

	It("returns the discount card for registered to the email", func() {
		respBody := `{
		  "cardNumber": "1234567891234567",
		  "discountPercentage": 15,
		  "issueNumber": 15,
		  "status": "VERIFIED"
		}`
		expectedCard := domain.Card{
			CardNumber:  "1234567891234567",
			IssueNumber: 15,
			Status:      "VERIFIED",
		}

		responseRecorder := httptest.NewRecorder()
		responseRecorder.Body = bytes.NewBufferString(respBody)
		resp := responseRecorder.Result()
		reqMatcher := checkoutsDiscountReqMatcher{
			host:  host,
			email: email,
		}
		mockHttpClient.EXPECT().
			Do(reqMatcher).
			Return(resp, nil)

		client := outbound.NewCheckoutsColleagueDiscountClient(host, mockHttpClient)
		card, err := client.GetDiscountCard(ctx, email)

		Expect(err).ToNot(HaveOccurred())
		Expect(card).To(Equal(expectedCard))
	})

})

type checkoutsDiscountReqMatcher struct {
	host  string
	email string
}

func (c checkoutsDiscountReqMatcher) Matches(x interface{}) bool {
	actualRequest := x.(*http.Request)
	methodMatch := Expect(actualRequest.Method).To(Equal(http.MethodGet))

	expectedUrl := fmt.Sprintf("https://%s/discount-card?email=%s", c.host, c.email)
	urlMatch := Expect(actualRequest.URL.String()).To(Equal(expectedUrl))

	return methodMatch && urlMatch
}
func (c checkoutsDiscountReqMatcher) String() string {
	return fmt.Sprintf("a get requet to checkouts colleague discount with the correct email %s", c.email)
}
