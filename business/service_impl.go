package business

import (
	"errors"
	"github.com/pashukhin/coins-test-task/repository"
	"github.com/pashukhin/coins-test-task/service"
)

type ServiceImpl interface {
	service.Service
	SetAccountRepository(accounts repository.AccountRepository)
	SetPaymentRepository(payments repository.PaymentRepository)
}

func NewService() ServiceImpl {
	return &serviceImpl{}
}

// constant errors
var errZeroAmount = errors.New("amount must be greater than 0")
var errNotEnoughBalance = errors.New("amount must be greater than sender balance")
var errDifferentCurrencies = errors.New("only transactions between accounts with the same currencies are allowed")

type serviceImpl struct {
	accounts repository.AccountRepository
	payments repository.PaymentRepository
}

func (s *serviceImpl) SetAccountRepository(accounts repository.AccountRepository) {
	s.accounts = accounts
}

func (s *serviceImpl) SetPaymentRepository(payments repository.PaymentRepository) {
	s.payments = payments
}