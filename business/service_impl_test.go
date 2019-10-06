package business

import (
	"errors"
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/repository"
	"reflect"
	"testing"
)

type accountRepositoryMock struct {
	data []*entity.Account
}

func (a *accountRepositoryMock) GetAll() (all []*entity.Account, err error) {
	return a.data, nil
}

func (a *accountRepositoryMock) Get(id int64) (account *entity.Account, err error) {
	for _, acc := range a.data {
		if acc.ID == int(id) {
			return acc, nil
		}
	}
	return nil, errors.New("not found")
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
		a.data = append(a.data, account)
	}
	return nil
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
			s := &serviceImpl{
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
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceImpl{
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
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Payment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceImpl{
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantP   *entity.Payment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceImpl{
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

func Test_serviceImpl_checkSend(t *testing.T) {
	type fields struct {
		accounts repository.AccountRepository
		payments repository.PaymentRepository
	}
	type args struct {
		fromID int64
		toID   int64
		amount float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantAccFrom *entity.Account
		wantAccTo   *entity.Account
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceImpl{
				accounts: tt.fields.accounts,
				payments: tt.fields.payments,
			}
			gotAccFrom, gotAccTo, err := s.checkSend(tt.args.fromID, tt.args.toID, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkSend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccFrom, tt.wantAccFrom) {
				t.Errorf("checkSend() gotAccFrom = %v, want %v", gotAccFrom, tt.wantAccFrom)
			}
			if !reflect.DeepEqual(gotAccTo, tt.wantAccTo) {
				t.Errorf("checkSend() gotAccTo = %v, want %v", gotAccTo, tt.wantAccTo)
			}
		})
	}
}