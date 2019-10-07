package saga

type Saga interface {
	Init(stages ...Stage) error
	Run() (err error)
}
