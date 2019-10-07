package stages

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

// DebitSourceStage if first stage of payment processing.
// Debits source account.
type DebitSourceStage struct {
	accounts repository.AccountRepository
	accFrom  *entity.Account
	amount   float64
	err      error
}

// NewDebitSourceStage makes new DebitSourceStage with parameters.
func NewDebitSourceStage(accounts repository.AccountRepository, accFrom *entity.Account, amount float64) *DebitSourceStage {
	return &DebitSourceStage{
		accounts: accounts,
		accFrom:  accFrom,
		amount:   amount,
	}
}

// Ahead decreases accFrom.Balance on amount using account repository.
func (s *DebitSourceStage) Ahead() error {
	s.err = s.accounts.Debit(s.accFrom, s.amount)
	return s.err
}

// Back increases accFrom.Balance on amount using account repository.
func (s *DebitSourceStage) Back() error {
	s.err = s.accounts.Credit(s.accFrom, s.amount)
	return s.err
}
