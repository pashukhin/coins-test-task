package business

import (
	"github.com/pashukhin/coins-test-task/business/stages"
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/saga"
)

// Send makes a payment between source and target accounts.
// It makes some checks, finds accounts, makes and start saga for payment processing.
func (s *Logic) Send(fromID, toID int64, amount float64) (p *entity.Payment, err error) {
	//check business logic
	var accFrom, accTo *entity.Account
	if accFrom, accTo, err = s.checkSend(fromID, toID, amount); err != nil {
		return
	}

	debitSourceStage := stages.NewDebitSourceStage(s.accounts, accFrom, amount)
	createPaymentStage := stages.NewCreatePaymentStage(s.payments, accFrom, accTo, amount)
	creditTargetStage := stages.NewCreditTargetStage(s.accounts, accTo, amount)

	paymentSaga := saga.NewSagaImpl()
	if err = paymentSaga.AddStages(debitSourceStage, createPaymentStage, creditTargetStage); err != nil {
		return
	}
	if err = paymentSaga.Run(); err != nil {
		return
	}
	p = createPaymentStage.GetResult()
	return
}
