package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Storage provides storage operations with accounts and payments
type Storage interface {
	// Accounts returns list of all accounts in db or error in case of fail.
	Accounts() ([]*Account, error)
	// GetAccountBy returns single account by column value from db or error in case of fail.
	GetAccountBy(column string, value interface{}) (*Account, error)
	// Payments returns list of all payments from db of error in case of fail.
	Payments() ([]*Payment, error)
	// Srore stores from and to accounts and payment to db in single transaction.
	// Returns newly created payment entity from db or error in any case of fail.
	// It mades rollback automatically in case of any fail.
	// Throw panic after rollback error
	Store(from, to *Account, payment *Payment) (*Payment, error)
}

type storage struct {
	db *sqlx.DB
}

func (r storage) Accounts() (all []*Account, err error) {
	err = r.db.Select(&all, "select * from account")
	return
}

func (r storage) GetAccountBy(column string, value interface{}) (a *Account, err error) {
	sql := fmt.Sprintf("select * from account where \"%s\" = $1", column)
	a = &Account{}
	err = r.db.Get(a, sql, value)
	return
}

func (r storage) Payments() (all []*Payment, err error) {
	err = r.db.Select(&all, "select * from payment")
	return
}

func (r storage) txRollback(tx *sqlx.Tx) {
	if err := tx.Rollback(); err != nil {
		panic(err)
	}
}

// txExec wraps sqlx.Tx.Exec with rollback on error
func (r storage) txExec(tx *sqlx.Tx, query string, args ...interface{}) (err error) {
	defer func() {
		if err != nil {
			r.txRollback(tx)
		}
	}()
	if _, err = tx.Exec(query, args...); err != nil {
		return
	}
	return
}

// txGet wraps sqlx.Tx.Get with rollback on error
func (r storage) txGet(tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) (err error) {
	defer func() {
		if err != nil {
			r.txRollback(tx)
		}
	}()
	if err = tx.Get(dest, query, args...); err != nil {
		return
	}
	return
}

func (r storage) Store(from, to *Account, payment *Payment) (p *Payment, err error) {
	var tx *sqlx.Tx
	if tx, err = r.db.Beginx(); err != nil {
		return
	}
	p = &Payment{}
	sql := "insert into payment (account_from_id, account_to_id, amount) values ($1, $2, $3) returning *"
	if err = r.txGet(tx, p, sql, from.ID, to.ID, payment.Amount); err != nil {
		return
	}
	if err = r.txExec(tx, "update account set balance = $1 where id = $2", from.Balance, from.ID); err != nil {
		return
	}
	if err = r.txExec(tx, "update account set balance = $1 where id = $2", to.Balance, to.ID); err != nil {
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	return
}
