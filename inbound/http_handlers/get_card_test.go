package http_handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	apiproblem "github.com/JSainsburyPLC/go-api-problem"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/inbound/http_handlers"
	mock_http_handlers "github.com/JSainsburyPLC/smartshop-api-colleague-discount/mocks/inbound/http_handlers"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("get card handler", func() {

	var (
		ctrl                               *gomock.Controller
		mockColleagueDiscountCardRetriever *mock_http_handlers.MockColleagueDiscountCardRetriever

		apiProblemFactory apiproblem.Factory
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockColleagueDiscountCardRetriever = mock_http_handlers.NewMockColleagueDiscountCardRetriever(ctrl)

		apiProblemFactory = apiproblem.NewFactory("http://www.github.com/JSainsburyPLC/smartshop-api-colleague-discount")
	})

	It("returns 200 and the card for the user", func() {
		userId := 1234
		usersCard := domain.Card{
			CardNumber:  "1234567891234567",
			IssueNumber: 15,
			Status:      "VERIFIED",
		}
		url := fmt.Sprintf("/discount-card?user_id=%d", userId)
		request, _ := http.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		mockColleagueDiscountCardRetriever.EXPECT().
			GetCardForUser(request.Context(), userId).
			Return(usersCard, nil)

		handler := http_handlers.NewGetCardHandler(apiProblemFactory, mockColleagueDiscountCardRetriever)
		handler.ServeHTTP(response, request)

		Expect(response.Code).To(Equal(http.StatusOK))
		expectedBody := `{
			"card_number": "1234567891234567",
			"issue_number": "15",
			"status": "VERIFIED"
		}`
		Expect(response.Body).To(MatchJSON(expectedBody))
	})

	DescribeTable("the user id is a not a positive integer",
		func(userId string) {
			url := fmt.Sprintf("/discount-card?user_id=%s", userId)
			request, _ := http.NewRequest(http.MethodGet, url, nil)
			response := httptest.NewRecorder()
			expectedBody := `{
  "type": "http://www.github.com/JSainsburyPLC/smartshop-api-colleague-discount",
  "status": 400,
  "title": "Bad Request",
  "detail": "user id must be a positive integer",
  "code": 39001
}`

			handler := http_handlers.NewGetCardHandler(apiProblemFactory, mockColleagueDiscountCardRetriever)
			handler.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(response.Body).To(MatchJSON(expectedBody))

		},
		Entry("when it is zero", "0"),
		Entry("when it is negative", "-1"),
		Entry("when it is a float", "1.1"),
		Entry("when it is not a number", "123-1232"),
	)
})
