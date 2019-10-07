package stages

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

// CreatePaymentStage is second stage of payment processing.
// Creates payment and stores it in result field
type CreatePaymentStage struct {
	payments       repository.PaymentRepository
	accFrom, accTo *entity.Account
	amount         float64
	result         *entity.Payment
	err            error
}

// NewCreatePaymentStage creates new CreatePaymentStage
// Returns *CreatePaymentStage.
func NewCreatePaymentStage(payments repository.PaymentRepository, accFrom, accTo *entity.Account, amount float64) *CreatePaymentStage {
	return &CreatePaymentStage{
		payments: payments,
		accFrom:  accFrom,
		accTo:    accTo,
		amount:   amount,
	}
}

// Ahead is implementation of Stage.Ahead.
// It just calls Create on payment repository and stores result.
// Returns error from repository.
func (s *CreatePaymentStage) Ahead() error {
	s.result, s.err = s.payments.Create(s.accFrom, s.accTo, s.amount)
	return s.err
}

// Back is implementation of Back.Ahead.
// It deletes previously created payment calling repository method Delete.
// Returns error from repository.
func (s *CreatePaymentStage) Back() error {
	s.err = s.payments.Delete(s.result)
	return s.err
}

// GetResult returns result makes on Ahead method
func (s *CreatePaymentStage) GetResult() *entity.Payment {
	return s.result
}
