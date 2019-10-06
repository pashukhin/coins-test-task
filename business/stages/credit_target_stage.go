package business

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

type creditTargetStage struct {
	accounts repository.AccountRepository
	accTo    *entity.Account
	amount   float64
	err      error
}

func (s *creditTargetStage) Ahead() error {
	s.err = s.accounts.Credit(s.accTo, s.amount)
	return s.err
}

func (s *creditTargetStage) Back() error {
	s.err = s.accounts.Debit(s.accTo, s.amount)
	return s.err
}
