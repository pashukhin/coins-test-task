package stages

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

type CreditTargetStage struct {
	accounts repository.AccountRepository
	accTo    *entity.Account
	amount   float64
	err      error
}

func NewCreditTargetStage(accounts repository.AccountRepository, accTo *entity.Account, amount float64) *CreditTargetStage {
	return &CreditTargetStage{
		accounts: accounts,
		accTo:    accTo,
		amount:   amount,
	}
}

func (s *CreditTargetStage) Ahead() error {
	s.err = s.accounts.Credit(s.accTo, s.amount)
	return s.err
}

func (s *CreditTargetStage) Back() error {
	s.err = s.accounts.Debit(s.accTo, s.amount)
	return s.err
}
