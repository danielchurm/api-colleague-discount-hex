package domain

import "context"

type Card struct {
	CardNumber  string
	IssueNumber int
	Status      string // Possible values: NEW, VERIFIED
}

//go:generate mockgen -destination=../mocks/domain/discount_cards.go -source=discount_cards.go

type UserRepository interface {
	GetEmail(ctx context.Context, userId int) (string, error)
}

type DiscountCardRepository interface {
	GetDiscountCard(ctx context.Context, email string) (Card, error)
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

func (dc DiscountCardRetriever) GetCardForUser(ctx context.Context, userId int) (Card, error) {
	email, err := dc.userRepo.GetEmail(ctx, userId)
	if err != nil {
		return Card{}, err
	}

	discountCard, err := dc.discountCardRepo.GetDiscountCard(ctx, email)
	if err != nil {
		return Card{}, err
	}

	return discountCard, nil
}
