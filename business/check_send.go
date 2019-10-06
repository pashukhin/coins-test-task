package business

import "github.com/pashukhin/coins-test-task/entity"

// checkSend checks business conditions for Send operation
// returns accounts or error if conditions are not met
func (s *serviceImpl) checkSend(fromID, toID int64, amount float64) (accFrom, accTo *entity.Account, err error) {
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
