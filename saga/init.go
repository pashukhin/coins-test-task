package saga

import "errors"

// AddStage adds stage into sagaImpl, like it described in its name
func (s *sagaImpl) Init(stages ...Stage) error {
	if len(stages) == 0 {
		return errors.New("empty stage list")
	}
	for _, stage := range stages {
		s.stages = append(s.stages, stage)
	}
	return nil
}
