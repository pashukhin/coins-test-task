package entity

// Account represents an user's account
type Account struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
