// Package middleware contains basic middleware interface and its implementations.
package middleware

import "github.com/pashukhin/coins-test-task/service"

// Middleware is common interface for application middlewares.
// It extends service.Service with SetNext method.
type Middleware interface {
	service.Service
	SetNext(service service.Service)
}

type middleware struct {
	next service.Service
}

// SetNext sets next service in chain of middlewares.
func (m *middleware) SetNext(service service.Service) {
	m.next = service
}
