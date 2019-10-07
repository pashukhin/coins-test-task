package business

import (
	"errors"
	"github.com/pashukhin/coins-test-task/entity"
)

// constant errors
var (
	errZeroAmount          = errors.New("amount must be greater than 0")
	errNotEnoughBalance    = errors.New("amount must be greater than sender balance")
	errDifferentCurrencies = errors.New("only transactions between accounts with the same currencies are allowed")
)

// checkSend checks business conditions for Send operation
// returns accounts or error if conditions are not met
func (s *Logic) checkSend(fromID, toID int64, amount float64) (accFrom, accTo *entity.Account, err error) {
	//check business logic
	// zero amount in request
	if amount <= 0 {
		err = errZeroAmount
		return
	}
	// is first account exists?
	if accFrom, err = s.accounts.Get(fromID); err != nil {
		return
	}
	// check amount of first account
	if accFrom.Balance < amount {
		err = errNotEnoughBalance
		return
	}
	// is second account exists?
	if accTo, err = s.accounts.Get(toID); err != nil {
		return
	}
	// are currencies of accounts equivalent?
	if accFrom.Currency != accTo.Currency {
		err = errDifferentCurrencies
		return
	}
	return
}
