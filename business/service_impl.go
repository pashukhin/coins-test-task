package business

import (
	"github.com/pashukhin/coins-test-task/entity"
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

func (s *serviceImpl) Account(id int64) (account *entity.Account, err error) {
	account, err = s.accounts.Get(id)
	return
}