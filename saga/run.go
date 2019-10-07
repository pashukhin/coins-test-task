package saga

// Run runs sagaImpl stage-by-stage.
// If any stage fails, sagaImpl runs back functions for successfully completed stages in reversed order
func (s *sagaImpl) Run() (err error) {
	l := len(s.stages)
	for c := 0; c < l; c++  {
		if err = s.stages[c].Ahead(); err != nil {
			for i := c-1; i >= 0; i++ {
				if err = s.stages[i].Back(); err != nil {
					return
				}
			}
			break
		}
	}
	return
}
