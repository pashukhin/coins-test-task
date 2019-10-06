package service

import "github.com/pashukhin/coins-test-task/entity"

// Service provides all operations.
type Service interface {
	// Accounts returns list of all accounts
	Accounts() ([]*entity.Account, error)
	// Payments returns list of all payments
	Payments() ([]*entity.Payment, error)
	// Send builds a payment from one account to another
	Send(fromID, toID int64, amount float64) (*entity.Payment, error)
	Account(id int64) (account *entity.Account, err error)
}
