package middleware

import "github.com/pashukhin/coins-test-task/service"

type Middleware interface {
	service.Service
	SetNext(service service.Service)
}

type middleware struct {
	next service.Service
}

func (m *middleware) SetNext(service service.Service) {
	m.next = service
}