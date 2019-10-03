package main

import (
	"time"

	"errors"
)

// Service provides all operations.
type Service interface {
	// Accounts returns list of all accounts
	Accounts() ([]*Account, error)
	// Payments returns list of all payments
	Payments() ([]*Payment, error)
	// Send builds a payment from one account to another
	Send(fromID, toID string, amount float64) (*Payment, error)
}

// Account represents an user's account
type Account struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

// Payment represents a payment between two account
type Payment struct {
	ID      int       `json:"id"`
	FromID  int       `json:"from_id" db:"account_from_id"`
	ToID    int       `json:"to_id" db:"account_to_id"`
	Created time.Time `json:"created"`
	Amount  float64   `json:"amount"`
}

type service struct {
	storage Storage
}

func (s service) Accounts() ([]*Account, error) {
	return s.storage.Accounts()
}

func (s service) Payments() ([]*Payment, error) {
	return s.storage.Payments()
}

func (s service) Send(fromID, toID string, amount float64) (*Payment, error) {
	if amount <= 0 {
		return nil, errZeroAmount
	}
	accFrom, err := s.storage.GetAccountBy("name", fromID)
	if err != nil {
		return nil, err
	}
	if accFrom.Balance < amount {
		return nil, errNotEnauthBalance
	}
	accTo, err := s.storage.GetAccountBy("name", toID)
	if err != nil {
		return nil, err
	}
	if accFrom.Currency != accTo.Currency {
		return nil, errDifferentCurrencies
	}
	accFrom.Balance -= amount
	accTo.Balance += amount
	payment := &Payment{
		FromID: accFrom.ID,
		ToID:   accTo.ID,
		Amount: amount,
	}
	return s.storage.Store(accFrom, accTo, payment)
}

var errZeroAmount = errors.New("amount must be greater than 0")
var errNotEnauthBalance = errors.New("amount must be greater than sender balance")
var errDifferentCurrencies = errors.New("only transactions between accounts with the same currencies are allowed")
