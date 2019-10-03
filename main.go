package main

import (
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"

	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	db, err := sqlx.Connect("postgres", "host=db port=5432 user=user password=password dbname=db sslmode=disable")
	if err != nil {
		logger.Log("err", err)
		return
	}

	var svc Service
	svc = service{
		storage: storage{db},
	}
	svc = loggingMiddleware{logger, svc}

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "group",
		Subsystem: "service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "group",
		Subsystem: "service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	svc = instrumentingMiddleware{requestCount, requestLatency, svc}

	accountsHandler := httptransport.NewServer(
		makeAccountsEndpoint(svc),
		decodeAccountsRequest,
		encodeResponse,
	)
	paymentsHandler := httptransport.NewServer(
		makePaymentsEndpoint(svc),
		decodePaymentsRequest,
		encodeResponse,
	)
	sendHandler := httptransport.NewServer(
		makeSendEndpoint(svc),
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
