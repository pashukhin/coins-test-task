package service

import "github.com/pashukhin/coins-test-task/entity"

// Service provides all operations.
type Service interface {
	Accounts() ([]*entity.Account, error)
	Payments() ([]*entity.Payment, error)
	Send(fromID, toID int64, amount float64) (*entity.Payment, error)
	Account(id int64) (account *entity.Account, err error)
}
