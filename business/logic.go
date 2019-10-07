package business

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
)

// NewLogic makes a new Logic and returns it as ServiceImpl
func NewLogic() ServiceImpl {
	return &Logic{}
}

// Logic implements business logic.
type Logic struct {
	accounts repository.AccountRepository
	payments repository.PaymentRepository
}

// SetAccountRepository makes only just described in its name.
// Used for inversion of control pattern.
func (s *Logic) SetAccountRepository(accounts repository.AccountRepository) {
	s.accounts = accounts
}

// SetPaymentRepository makes only just described in its name.
// Used for inversion of control pattern.
func (s *Logic) SetPaymentRepository(payments repository.PaymentRepository) {
	s.payments = payments
}

// Account returns an account by its id or error if fails.
func (s *Logic) Account(id int64) (account *entity.Account, err error) {
	account, err = s.accounts.Get(id)
	return
}
