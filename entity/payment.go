package entity

import "time"

// Payment represents a payment between two account
type Payment struct {
	ID      int       `json:"id"`
	FromID  int       `json:"from_id" db:"account_from_id"`
	ToID    int       `json:"to_id" db:"account_to_id"`
	Created time.Time `json:"created"`
	Amount  float64   `json:"amount"`
}
