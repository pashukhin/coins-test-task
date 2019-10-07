package saga

// sagaImpl is very simple implementation of sagaImpl pattern
// see https://en.wikipedia.org/wiki/Compensating_transaction
type sagaImpl struct {
	// list of stages to apply
	stages []Stage
}

// NewSagaImpl makes a new sagaImpl, like it described in its name
func NewSagaImpl() Saga {
	return &sagaImpl{stages: make([]Stage, 0)}
}
