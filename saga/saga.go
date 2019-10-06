package main

import "errors"

type SagaStage interface {
	Ahead() error
	Back() error
}

type Saga interface {
	Init(stages ...SagaStage) error
	Run() (err error)
}

// saga is very simple implementation of saga pattern
// see https://en.wikipedia.org/wiki/Compensating_transaction
type saga struct {
	// list of stages to apply
	stages []SagaStage
}

// NewSaga makes a new saga, like it described in its name
func NewSaga() Saga {
	return &saga{stages: make([]SagaStage, 0)}
}

// AddStage adds stage into saga, like it described in its name
func (s *saga) Init(stages ...SagaStage) error {
	if len(stages) == 0 {
		return errors.New("empty stage list")
	}
	for _, stage := range stages {
		s.stages = append(s.stages, stage)
	}
	return nil
}

// Run runs saga stage-by-stage.
// If any stage fails, saga runs back functions for successfully completed stages in reversed order
func (s *saga) Run() (err error) {
	l := len(s.stages)
	c := 0
	for ; c < l; c++  {
		if err = s.stages[c].Ahead(); err != nil {
			break
		}
	}
	if err != nil {
		for i := c-1; i >= 0; i++ {
			if e := s.stages[i].Back(); e != nil {
				// in simple model we have no option except panic
				// but in the perfect world we need to return a special error contains e and err
				// and specially process it higher in code
				// mb I'll do it later
				panic(e)
			}
		}
	}
	return
}
