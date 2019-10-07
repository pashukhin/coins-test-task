package saga

// SimpleSaga is very simple implementation of Saga.
type SimpleSaga struct {
	// list of stages to apply
	stages []Stage
}

// NewSagaImpl makes a new SimpleSaga, like it described in its name
func NewSagaImpl() Saga {
	return &SimpleSaga{stages: make([]Stage, 0)}
}
