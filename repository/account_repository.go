package main

type AccountRepository interface {
	GetAll() (all []*Account, err error)
	Get(id int64) (account *Account, err error)
	Debit(acc *Account, amount float64) error
	Credit(acc *Account, amount float64) error
}

type accountRepository struct {
	repository
}

func (a accountRepository) GetAll() (all []*Account, err error) {
	err = a.db.Select(&all, "select * from account")
	return
}

func (a accountRepository) Get(id int64) (account *Account, err error) {
	account = &Account{}
	err = a.db.Get(account, "select * from account where id = $1", id)
	return
}


func (a accountRepository) Debit(acc *Account, amount float64) error {
	return a.ExecForOne("update account set balance = balance - $1 where id = $2 and balance >= $1", amount, acc.ID)
}

func (a accountRepository) Credit(acc *Account, amount float64) error {
	return a.ExecForOne("update account set balance = balance + $1 where id = $2", amount, acc.ID)
}



