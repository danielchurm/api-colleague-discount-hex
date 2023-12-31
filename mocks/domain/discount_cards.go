// Code generated by MockGen. DO NOT EDIT.
// Source: discount_cards.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	domain "github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetEmail mocks base method.
func (m *MockUserRepository) GetEmail(ctx context.Context, userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail", ctx, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockUserRepositoryMockRecorder) GetEmail(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockUserRepository)(nil).GetEmail), ctx, userId)
}

// MockDiscountCardRepository is a mock of DiscountCardRepository interface.
type MockDiscountCardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDiscountCardRepositoryMockRecorder
}

// MockDiscountCardRepositoryMockRecorder is the mock recorder for MockDiscountCardRepository.
type MockDiscountCardRepositoryMockRecorder struct {
	mock *MockDiscountCardRepository
}

// NewMockDiscountCardRepository creates a new mock instance.
func NewMockDiscountCardRepository(ctrl *gomock.Controller) *MockDiscountCardRepository {
	mock := &MockDiscountCardRepository{ctrl: ctrl}
	mock.recorder = &MockDiscountCardRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiscountCardRepository) EXPECT() *MockDiscountCardRepositoryMockRecorder {
	return m.recorder
}

// GetDiscountCard mocks base method.
func (m *MockDiscountCardRepository) GetDiscountCard(ctx context.Context, email string) (domain.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiscountCard", ctx, email)
	ret0, _ := ret[0].(domain.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiscountCard indicates an expected call of GetDiscountCard.
func (mr *MockDiscountCardRepositoryMockRecorder) GetDiscountCard(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiscountCard", reflect.TypeOf((*MockDiscountCardRepository)(nil).GetDiscountCard), ctx, email)
}
