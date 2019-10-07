// Package saga describes interface for saga - procedure for making consistent non-transactional operations in heterogeneous environment.
// About saga pattern see https://en.wikipedia.org/wiki/Compensating_transaction
package saga

// Saga interface declares methods for saga implementations.
type Saga interface {
	AddStages(stages ...Stage) error
	Run() (err error)
}
