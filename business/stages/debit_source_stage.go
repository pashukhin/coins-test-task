package stages

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

type DebitSourceStage struct {
	accounts repository.AccountRepository
	accFrom  *entity.Account
	amount   float64
	err      error
}

func NewDebitSourceStage(accounts repository.AccountRepository, accFrom *entity.Account, amount float64) *DebitSourceStage {
	return &DebitSourceStage{
		accounts: accounts,
		accFrom:  accFrom,
		amount:   amount,
	}
}

func (s *DebitSourceStage) Ahead() error {
	s.err = s.accounts.Debit(s.accFrom, s.amount)
	return s.err
}

func (s *DebitSourceStage) Back() error {
	s.err = s.accounts.Credit(s.accFrom, s.amount)
	return s.err
}
