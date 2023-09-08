package outbound_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	mock_outbound "github.com/JSainsburyPLC/smartshop-api-colleague-discount/mocks/outbound"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/outbound"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOutboundSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Outbound Test Suite")
}

var _ = Describe("identity orchestrator client", func() {
	var (
		ctrl           *gomock.Controller
		mockHttpClient *mock_outbound.MockHttpClient

		host   = "api-identity-orchestrator-ecs.app.internal"
		apikey = "gfdg-545665-fgdfg"
		reqId  = "g54df-g5df4gd-dfg54-dfg4"
		userId = "123"

		ctx context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockHttpClient = mock_outbound.NewMockHttpClient(ctrl)

		ctx = context.Background()
		ctx = context.WithValue(ctx, log.ContextKeyRequestID, reqId)
		ctx = context.WithValue(ctx, log.ContextKeyUserID, userId)

	})

	It("returns the email address of the user", func() {
		userId := 123
		expectedEmail := "user123@example.com"
		respBody := `{
			"id": 1,
			"identity_uuid": "",
			"email": "user123@example.com",
			"title": "",
			"first_name": "Johnny",
			"last_name": "Bravo",
			"terms_accepted": "0001-01-01T00:00:00Z",
			"nectar_card": "111222333444",
			"marketing_emails_preference": false,
			"registration_date": "0001-01-01T00:00:00Z",
			"is_guest_user": false,
			"confirmed_account": false,
			"has_migration_just_occurred": false,
			"is_locked": false,
			"last_login": "2021-01-01T00:00:00Z",
			"last_edited": "2021-01-01T01:00:00Z"
		}`

		responseRecorder := httptest.NewRecorder()
		responseRecorder.Write([]byte(respBody))
		resp := responseRecorder.Result()
		reqMatcher := identityOrchReqMatcher{
			host:   host,
			userId: userId,
			apikey: apikey,
			reqId:  reqId,
		}
		mockHttpClient.EXPECT().Do(reqMatcher).Return(resp, nil)

		client := outbound.NewIdentityOrchestratorClient(host, apikey, mockHttpClient)
		email, err := client.GetEmail(ctx, userId)

		Expect(err).ToNot(HaveOccurred())
		Expect(email).To(Equal(expectedEmail))
	})
})

type identityOrchReqMatcher struct {
	host   string
	userId int
	apikey string
	reqId  string
}

func (m identityOrchReqMatcher) Matches(x interface{}) bool {
	actualRequest := x.(*http.Request)
	methodMatch := Expect(actualRequest.Method).To(Equal(http.MethodGet))

	expectedUrl := fmt.Sprintf("http://%s/api/v1/users/%d", m.host, m.userId)
	urlMatch := Expect(actualRequest.URL.String()).To(Equal(expectedUrl))

	reqIdMatch := Expect(actualRequest.Header.Get("X-Request-Id")).To(Equal(m.reqId))
	userIdMatch := Expect(actualRequest.Header.Get("X-User")).To(Equal(strconv.Itoa(m.userId)))
	apikeyMatch := Expect(actualRequest.Header.Get("X-API-KEY")).To(Equal(m.apikey))

	return methodMatch && urlMatch && reqIdMatch && userIdMatch && apikeyMatch
}
func (m identityOrchReqMatcher) String() string {
	return fmt.Sprintf("a get requet to identity orchestrator for user %d with correct headers", m.userId)
}
