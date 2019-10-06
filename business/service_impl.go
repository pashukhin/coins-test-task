package business

import (
	"errors"
	"github.com/pashukhin/coins-test-task/business/stages"
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
	"github.com/pashukhin/coins-test-task/saga"
	"github.com/pashukhin/coins-test-task/service"
)

type ServiceImpl interface {
	service.Service
	SetAccountRepository(accounts repository.AccountRepository)
	SetPaymentRepository(payments repository.PaymentRepository)
}

func NewService() ServiceImpl {
	return &serviceImpl{}
}

// constant errors
var errZeroAmount = errors.New("amount must be greater than 0")
var errNotEnoughBalance = errors.New("amount must be greater than sender balance")
var errDifferentCurrencies = errors.New("only transactions between accounts with the same currencies are allowed")

type serviceImpl struct {
	accounts repository.AccountRepository
	payments repository.PaymentRepository
}

func (s *serviceImpl) SetAccountRepository(accounts repository.AccountRepository) {
	s.accounts = accounts
}

func (s *serviceImpl) SetPaymentRepository(payments repository.PaymentRepository) {
	s.payments = payments
}

func (s *serviceImpl) Send(fromID, toID int64, amount float64) (p *entity.Payment, err error) {
	//check business logic
	var accFrom, accTo *entity.Account
	if accFrom, accTo, err = s.checkSend(fromID, toID, amount); err != nil {
		return
	}

	debitSourceStage := stages.NewDebitSourceStage(s.accounts, accFrom, amount)
	createPaymentStage := stages.NewCreatePaymentStage(s.payments, accFrom, accTo, amount)
	creditTargetStage := stages.NewCreditTargetStage(s.accounts, accTo, amount)

	// make new and very local saga
	paymentSaga := saga.NewSaga()
	if err = paymentSaga.Init(debitSourceStage, createPaymentStage, creditTargetStage); err != nil{
		return
	}
	if err = paymentSaga.Run(); err != nil{
		return
	}
	p = createPaymentStage.GetResult()
	return
}