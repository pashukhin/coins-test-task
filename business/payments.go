package business

import "github.com/pashukhin/coins-test-task/entity"

func (s *serviceImpl) Payments() ([]*entity.Payment, error) {
	return s.payments.GetAll()
}
