package domain_test

import (
	"context"
	"testing"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
	mock_domain "github.com/JSainsburyPLC/smartshop-api-colleague-discount/mocks/domain"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDomainSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Test Suite")
}

var _ = Describe("discount cards", func() {

	var (
		ctrl                 *gomock.Controller
		mockUserRepo         *mock_domain.MockUserRepository
		mockDiscountCardRepo *mock_domain.MockDiscountCardRepository

		discountCardRetriever domain.DiscountCardRetriever
		ctx                   context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockUserRepo = mock_domain.NewMockUserRepository(ctrl)
		mockDiscountCardRepo = mock_domain.NewMockDiscountCardRepository(ctrl)

		discountCardRetriever = domain.NewDiscountCard(mockUserRepo, mockDiscountCardRepo)
		ctx = context.Background()
	})

	It("returns the user's discount card", func() {
		userId := 123
		email := "user123@example.com"
		expectedCard := domain.Card{
			CardNumber:  "1234567891234567",
			IssueNumber: 15,
			Status:      "VERIFIED",
		}

		mockUserRepo.EXPECT().GetEmail(ctx, userId).Return(email, nil)
		mockDiscountCardRepo.EXPECT().GetDiscountCard(ctx, email).Return(expectedCard, nil)

		card, err := discountCardRetriever.GetCardForUser(ctx, userId)

		Expect(err).ToNot(HaveOccurred())
		Expect(card).To(Equal(expectedCard))
	})
})
