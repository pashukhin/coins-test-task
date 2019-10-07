package saga

// Stage is interface for saga stages
type Stage interface {
	// Ahead makes stage actions
	Ahead() error

	// Back cancels actions done with Ahead
	Back() error
}
