package business

import (
	"github.com/pashukhin/coins-test-task/repository"
	"github.com/pashukhin/coins-test-task/service"
)

// ServiceImpl is interface for business logic implementations.
// It adds two methods for set entity repositories to Service
type ServiceImpl interface {
	service.Service
	SetAccountRepository(accounts repository.AccountRepository)
	SetPaymentRepository(payments repository.PaymentRepository)
}
