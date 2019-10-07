package saga

type Stage interface {
	Ahead() error
	Back() error
}
