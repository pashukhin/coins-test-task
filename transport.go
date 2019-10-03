package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Accounts()
		if err != nil {
			return accountsResponse{v, err.Error()}, nil
		}
		return accountsResponse{v, ""}, nil
	}
}

func decodeAccountsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request accountsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type accountsRequest struct {
	// techincally, here may be some filters or so
}

type accountsResponse struct {
	Accounts []*Account `json:"accounts"`
	Err      string     `json:"err,omitempty"`
}

func makePaymentsEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Payments()
		if err != nil {
			return paymentsResponse{v, err.Error()}, nil
		}
		return paymentsResponse{v, ""}, nil
	}
}

func decodePaymentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request paymentsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type paymentsRequest struct {
	// techincally, here may be some filters or so
}

type paymentsResponse struct {
	Payments []*Payment `json:"payments"`
	Err      string     `json:"err,omitempty"`
}

func makeSendEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(sendRequest)
		payment, err := svc.Send(req.From, req.To, req.Amount)
		if err != nil {
			return sendResponse{nil, err.Error()}, nil
		}
		return sendResponse{payment, ""}, nil
	}
}

func decodeSendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sendRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type sendRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	// techincally, here may be some filters or so
}

type sendResponse struct {
	Payment *Payment `json:"payment"`
	Err     string   `json:"err,omitempty"`
}
