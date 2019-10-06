package stages

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

type CreatePaymentStage struct {
	payments       repository.PaymentRepository
	accFrom, accTo *entity.Account
	amount         float64
	result         *entity.Payment
	err            error
}

func NewCreatePaymentStage(payments repository.PaymentRepository, accFrom, accTo *entity.Account, amount float64) *CreatePaymentStage {
	return &CreatePaymentStage{
		payments: payments,
		accFrom:  accFrom,
		accTo:    accTo,
		amount:   amount,
	}
}

func (s *CreatePaymentStage) Ahead() error {
	s.result, s.err = s.payments.Create(s.accFrom, s.accTo, s.amount)
	return s.err
}

func (s *CreatePaymentStage) Back() error {
	s.err = s.payments.Delete(s.result)
	return s.err
}

func (s *CreatePaymentStage) GetResult() *entity.Payment {
	return s.result
}
