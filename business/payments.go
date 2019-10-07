package business

import "github.com/pashukhin/coins-test-task/entity"

// Payments returns list of payments of error if fails.
func (s *Logic) Payments() ([]*entity.Payment, error) {
	return s.payments.GetAll()
}
