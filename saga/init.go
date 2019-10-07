package saga

import "errors"

// AddStages adds stage into SimpleSaga, like it described in its name
func (s *SimpleSaga) AddStages(stages ...Stage) error {
	if len(stages) == 0 {
		return errors.New("empty stage list")
	}
	for _, stage := range stages {
		s.stages = append(s.stages, stage)
	}
	return nil
}
