package business

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

type createPaymentStage struct {
	payments       repository.PaymentRepository
	accFrom, accTo *entity.Account
	amount         float64
	result         *entity.Payment
	err            error
}

func (s *createPaymentStage) Ahead() error {
	s.result, s.err = s.payments.Create(s.accFrom, s.accTo, s.amount)
	return s.err
}

func (s *createPaymentStage) Back() error {
	s.err = s.payments.Delete(s.result)
	return s.err
}
