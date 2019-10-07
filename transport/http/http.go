package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/pashukhin/coins-test-task/service"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
// Useful in a profilesvc server.
func MakeHTTPHandler(s service.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET	/account/						list accounts
	// GET	/payment/						list all payments
	// POST	/payment						create new payment
	// GET	/account/{id}					get single account
	// GET	/account/{id}/payment			list all payments from/to account
	// GET	/account/{id}/payment/incoming	list incoming payments for account
	// GET	/account/{id}/payment/outgoing	list incoming payments for account
	// GET	/payment/{id}					get single payment

	r.Methods("GET").Path("/account").Handler(httptransport.NewServer(
		e.GetAccounts,
		decodeGetAccountsRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/payment").Handler(httptransport.NewServer(
		e.GetPayments,
		decodeGetPaymentsRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/payment").Handler(httptransport.NewServer(
		e.PostPayment,
		decodePostPaymentRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/account/{id}").Handler(httptransport.NewServer(
		e.GetAccountEndpoint,
		decodeGetAccountRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetAccountsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return getAccountsRequest{}, nil
}

func decodeGetPaymentsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return getPaymentsRequest{}, nil
}

func decodePostPaymentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req postPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetAccountRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	strId, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.ParseInt(strId, 10, 0)
	if err != nil {
		return nil, err
	}
	return getAccountRequest{ID: id}, nil
}

type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case mux.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
