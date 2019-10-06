package main

import (
	"github.com/pashukhin/coins-test-task/business"
	"github.com/pashukhin/coins-test-task/middleware"
	"github.com/pashukhin/coins-test-task/repository"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	//db, err := sqlx.Connect("postgres", "host=db port=5432 user=user password=password dbname=db sslmode=disable")
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5432 user=user password=password dbname=db sslmode=disable")
	if err != nil {
		if err := logger.Log("err", err); err != nil {
			panic(err)
		}
		return
	}

	accounts := repository.NewAccountRepository(db)
	payments := repository.NewPaymentRepository(db)
	service := business.NewService()
	service.SetAccountRepository(accounts)
	service.SetPaymentRepository(payments)

	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	loggingMiddleware.SetNext(service)

	instrumentingMiddleware := middleware.NewInstrumentingMiddleware()
	instrumentingMiddleware.SetNext(loggingMiddleware)

	accountsHandler := httptransport.NewServer(
		makeAccountsEndpoint(instrumentingMiddleware),
		decodeAccountsRequest,
		encodeResponse,
	)
	paymentsHandler := httptransport.NewServer(
		makePaymentsEndpoint(instrumentingMiddleware),
		decodePaymentsRequest,
		encodeResponse,
	)
	sendHandler := httptransport.NewServer(
		makeSendEndpoint(instrumentingMiddleware),
		decodeSendRequest,
		encodeResponse,
	)

	http.Handle("/accounts", accountsHandler)
	http.Handle("/payments", paymentsHandler)
	http.Handle("/send", sendHandler)

	http.Handle("/metrics", promhttp.Handler())

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
