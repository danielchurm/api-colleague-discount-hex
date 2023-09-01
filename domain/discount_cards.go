package domain

type Card struct {
	CardNumber  string
	IssueNumber int
	Status      string // Possible values: NEW, VERIFIED
}

//go:generate mockgen -destination=../mocks/domain/discount_cards.go -source=discount_cards.go

type UserRepository interface {
	GetEmail(userId int) (string, error)
}

type DiscountCardRepository interface {
	GetDiscountCard(email string) (Card, error)
}

type DiscountCardRetriever struct {
	userRepo         UserRepository
	discountCardRepo DiscountCardRepository
}

func NewDiscountCard(userRepository UserRepository, discountCardRepository DiscountCardRepository) DiscountCardRetriever {
	return DiscountCardRetriever{
		userRepo:         userRepository,
		discountCardRepo: discountCardRepository,
	}
}

func (dc DiscountCardRetriever) GetCardForUser(userId int) (Card, error) {
	email, err := dc.userRepo.GetEmail(userId)
	if err != nil {
		return Card{}, err
	}

	discountCard, err := dc.discountCardRepo.GetDiscountCard(email)
	if err != nil {
		return Card{}, err
	}

	return discountCard, nil
}
