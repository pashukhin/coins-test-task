package business

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

type debitSourceStage struct {
	accounts repository.AccountRepository
	accFrom  *entity.Account
	amount   float64
	err      error
}

func (s *debitSourceStage) Ahead() error {
	s.err = s.accounts.Debit(s.accFrom, s.amount)
	return s.err
}

func (s *debitSourceStage) Back() error {
	s.err = s.accounts.Credit(s.accFrom, s.amount)
	return s.err
}
