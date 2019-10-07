package http

import (
	"context"
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAccounts endpoint.Endpoint
	GetPayments endpoint.Endpoint
	PostPayment endpoint.Endpoint
	GetAccountEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetAccounts: MakeGetAccountsEndpoint(s),
		GetPayments: MakeGetPaymentsEndpoint(s),
		PostPayment: MakePostPaymentEndpoint(s),
		GetAccountEndpoint: MakeGetAccountEndpoint(s),
	}
}

type getAccountsRequest struct {
	// mb some filters here
}

type getAccountsResponse struct {
	Accounts []*entity.Account `json:"accounts"`
	Err      error             `json:"err,omitempty"`
}

func MakeGetAccountsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req := request.(getAccountsRequest)
		accounts, e := s.Accounts()
		return getAccountsResponse{accounts, e}, e
	}
}

type getPaymentsRequest struct {
	// mb some filters here
}

type getPaymentsResponse struct {
	Payments []*entity.Payment `json:"payments"`
	Err      error             `json:"err,omitempty"`
}

func MakeGetPaymentsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req := request.(getAccounts)
		payments, e := s.Payments()
		return getPaymentsResponse{payments, e}, e
	}
}

type postPaymentRequest struct {
	From   int64   `json:"from"`
	To     int64   `json:"to"`
	Amount float64 `json:"amount"`
}

type postPaymentResponse struct {
	Payment *entity.Payment `json:"payment"`
	Err     error           `json:"err,omitempty"`
}

func MakePostPaymentEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postPaymentRequest)
		payment, err := s.Send(req.From, req.To, req.Amount)
		return postPaymentResponse{payment, err}, err
	}
}

type getAccountRequest struct {
	ID int64 `json:"id"`
}

type getAccountResponse struct {
	Accounts *entity.Account `json:"account"`
	Err      error             `json:"err,omitempty"`
}

func MakeGetAccountEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getAccountRequest)
		account, e := s.Account(req.ID)
		return getAccountResponse{account, e}, e
	}
}
