package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pashukhin/coins-test-task/entity"
)

// AccountRepository declares methods for account repository implementations.
type AccountRepository interface {
	GetAll() (all []*entity.Account, err error)
	Get(id int64) (account *entity.Account, err error)
	Debit(acc *entity.Account, amount float64) error
	Credit(acc *entity.Account, amount float64) error
}

// NewAccountRepository makes accountRepository and returns it as AccountRepository
func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepository{&repository{db}}
}

type accountRepository struct {
	*repository
}

func (a accountRepository) GetAll() (all []*entity.Account, err error) {
	err = a.db.Select(&all, "select * from account")
	return
}

func (a accountRepository) Get(id int64) (account *entity.Account, err error) {
	account = &entity.Account{}
	err = a.db.Get(account, "select * from account where id = $1", id)
	return
}

func (a accountRepository) Debit(acc *entity.Account, amount float64) error {
	return a.ExecForOne("update account set balance = balance - $1 where id = $2 and balance >= $1", amount, acc.ID)
}

func (a accountRepository) Credit(acc *entity.Account, amount float64) error {
	return a.ExecForOne("update account set balance = balance + $1 where id = $2", amount, acc.ID)
}
