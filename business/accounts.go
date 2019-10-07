package business

import "github.com/pashukhin/coins-test-task/entity"

// Accounts implements service.Accounts.
// Returns list of all accounts or error if fails
func (s *serviceImpl) Accounts() ([]*entity.Account, error) {
	return s.accounts.GetAll()
}
