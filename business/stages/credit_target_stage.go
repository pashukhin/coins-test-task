package stages

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

// CreditTargetStage is third stage of payment processing.
// Increases accTo.Balance to amount.
type CreditTargetStage struct {
	accounts repository.AccountRepository
	accTo    *entity.Account
	amount   float64
	err      error
}

// NewCreditTargetStage makes new CreditTargetStage.
func NewCreditTargetStage(accounts repository.AccountRepository, accTo *entity.Account, amount float64) *CreditTargetStage {
	return &CreditTargetStage{
		accounts: accounts,
		accTo:    accTo,
		amount:   amount,
	}
}

// Ahead credits target account.
func (s *CreditTargetStage) Ahead() error {
	s.err = s.accounts.Credit(s.accTo, s.amount)
	return s.err
}

// Back debits previously credited account.
func (s *CreditTargetStage) Back() error {
	s.err = s.accounts.Debit(s.accTo, s.amount)
	return s.err
}
