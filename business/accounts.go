package business

import "github.com/pashukhin/coins-test-task/entity"

func (s *serviceImpl) Accounts() ([]*entity.Account, error) {
	return s.accounts.GetAll()
}
