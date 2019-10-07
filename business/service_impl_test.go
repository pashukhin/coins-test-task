package business

import (
	"errors"
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
	"reflect"
	"testing"
)

type accountRepositoryMock struct {
	Data []*entity.Account
}

func (a *accountRepositoryMock) GetAll() (all []*entity.Account, err error) {
	return a.Data, nil
}

func (a *accountRepositoryMock) Get(id int64) (account *entity.Account, err error) {
	for _, acc := range a.Data {
		if acc.ID == id {
			account = acc
			return
		}
	}
	err = errors.New("not found")
	return
}

func (a *accountRepositoryMock) Debit(account *entity.Account, amount float64) error {
	acc, err := a.Get(int64(account.ID))
	if err != nil {
		return err
	}
	if acc.Balance < amount {
		return repository.ErrStorageRowsAffected
	}
	acc.Balance -= amount
	return nil
}

func (a *accountRepositoryMock) Credit(account *entity.Account, amount float64) error {
	acc, err := a.Get(int64(account.ID))
	if err != nil {
		return err
	}
	acc.Balance += amount
	return nil
}

func (a *accountRepositoryMock) add(account *entity.Account) error {
	_, err := a.Get(int64(account.ID))
	if err != nil {
		a.Data = append(a.Data, account)
	}
	return nil
}

type paymentRepositoryMock struct {
	Data []*entity.Payment
}

func (p *paymentRepositoryMock) GetAll() (all []*entity.Payment, err error) {
	all = p.Data
	return
}

func (p *paymentRepositoryMock) GetIncomingFor(acc *entity.Account) (list []*entity.Payment, err error) {
	list = make([]*entity.Payment, 0)
	for _, pm := range p.Data {
		if pm.ToID == acc.ID {
			list = append(list, pm)
		}
	}
	return
}

func (p *paymentRepositoryMock) GetOutgoingFor(acc *entity.Account) (list []*entity.Payment, err error) {
	list = make([]*entity.Payment, 0)
	for _, pm := range p.Data {
		if pm.FromID == acc.ID {
			list = append(list, pm)
		}
	}
	return
}

func (p *paymentRepositoryMock) GetAllFor(acc *entity.Account) (list []*entity.Payment, err error) {
	list = make([]*entity.Payment, 0)
	for _, pm := range p.Data {
		if pm.FromID == acc.ID || pm.ToID == acc.ID {
			list = append(list, pm)
		}
	}
	return
}

func (p *paymentRepositoryMock) Get(id int64) (payment *entity.Payment, err error) {
	for _, pm := range p.Data {
		if pm.ID == id {
			payment = pm
			return
		}
	}
	err = errors.New("not found")
	return
}

func (p *paymentRepositoryMock) Create(from, to *entity.Account, amount float64) (payment *entity.Payment, err error) {
	payment = &entity.Payment{ID: int64(len(p.Data)) + 1, FromID: from.ID, ToID: to.ID, Amount: amount}
	p.Data = append(p.Data, payment)
	return
}

func (p *paymentRepositoryMock) Delete(payment *entity.Payment) error {
	for i, pm := range p.Data {
		if pm.ID == payment.ID {
			p.Data = append(p.Data[:i], p.Data[:i+1]...)
			return nil
		}
	}
	return repository.ErrStorageRowsAffected
}

func Test_serviceImpl_Account(t *testing.T) {
	type fields struct {
		accounts repository.AccountRepository
		payments repository.PaymentRepository
	}
	type args struct {
		id int64
	}
	account := &entity.Account{ID: 1}
	accountRepository := &accountRepositoryMock{[]*entity.Account{account}}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantAccount *entity.Account
		wantErr     bool
	}{
		{
			"get existing",
			fields{
				accounts: accountRepository,
			},
			args{
				id: 1,
			},
			account,
			false,
		},
		{
			"get not existing",
			fields{
				accounts: accountRepository,
			},
			args{
				id: 2,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Logic{
				accounts: tt.fields.accounts,
				payments: tt.fields.payments,
			}
			gotAccount, err := s.Account(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Account() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("Account() gotAccount = %v, want %v", gotAccount, tt.wantAccount)
			}
		})
	}
}

func Test_serviceImpl_Accounts(t *testing.T) {
	type fields struct {
		accounts repository.AccountRepository
		payments repository.PaymentRepository
	}
	account := &entity.Account{ID: 1}
	accountRepository := &accountRepositoryMock{[]*entity.Account{account}}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Account
		wantErr bool
	}{
		{
			"default",
			fields{
				accounts: accountRepository,
			},
			accountRepository.Data,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Logic{
				accounts: tt.fields.accounts,
				payments: tt.fields.payments,
			}
			got, err := s.Accounts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Accounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Accounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceImpl_Payments(t *testing.T) {
	type fields struct {
		accounts repository.AccountRepository
		payments repository.PaymentRepository
	}
	payment := &entity.Payment{ID: 1}
	paymentRepository := &paymentRepositoryMock{[]*entity.Payment{payment}}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Payment
		wantErr bool
	}{
		{
			"default",
			fields{
				payments: paymentRepository,
			},
			paymentRepository.Data,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Logic{
				accounts: tt.fields.accounts,
				payments: tt.fields.payments,
			}
			got, err := s.Payments()
			if (err != nil) != tt.wantErr {
				t.Errorf("Payments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceImpl_Send(t *testing.T) {
	type fields struct {
		accounts repository.AccountRepository
		payments repository.PaymentRepository
	}
	type args struct {
		fromID int64
		toID   int64
		amount float64
	}
	account1, account2 := &entity.Account{ID: 1, Balance: 10}, &entity.Account{ID: 2, Balance: 0}
	account3 := &entity.Account{ID: 3, Balance: 10, Currency: "FRK"}
	accountRepository := &accountRepositoryMock{[]*entity.Account{account1, account2, account3}}
	paymentRepository := &paymentRepositoryMock{[]*entity.Payment{}}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantP   *entity.Payment
		wantErr bool
	}{
		{
			"default",
			fields{
				accountRepository,
				paymentRepository,
			},
			args{
				account1.ID,
				account2.ID,
				1,
			},
			&entity.Payment{ID: 1, FromID: account1.ID, ToID: account2.ID, Amount: 1},
			false,
		},
		{
			"zero amount",
			fields{
				accountRepository,
				paymentRepository,
			},
			args{
				account1.ID,
				account2.ID,
				0,
			},
			nil,
			true,
		},
		{
			"different currencies",
			fields{
				accountRepository,
				paymentRepository,
			},
			args{
				account1.ID,
				account3.ID,
				1,
			},
			nil,
			true,
		},
		{
			"amount too large",
			fields{
				accountRepository,
				paymentRepository,
			},
			args{
				account1.ID,
				account2.ID,
				100,
			},
			nil,
			true,
		},
		{
			"sender not exists",
			fields{
				accountRepository,
				paymentRepository,
			},
			args{
				100,
				account2.ID,
				100,
			},
			nil,
			true,
		},
		{
			"recipient not exists",
			fields{
				accountRepository,
				paymentRepository,
			},
			args{
				account1.ID,
				100,
				100,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Logic{
				accounts: tt.fields.accounts,
				payments: tt.fields.payments,
			}
			gotP, err := s.Send(tt.args.fromID, tt.args.toID, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("Send() gotP = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
