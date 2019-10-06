package main

import (
	"flag"
	"fmt"
	httpTransport "github.com/pashukhin/coins-test-task/transport/http"

	"github.com/pashukhin/coins-test-task/business"
	"github.com/pashukhin/coins-test-task/middleware"
	"github.com/pashukhin/coins-test-task/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)

	//db, err := sqlx.Connect("postgres", "host=db port=5432 user=user password=password dbname=db sslmode=disable")
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5432 user=user password=password dbname=db sslmode=disable")
	if err != nil {
		if e := logger.Log("err", err); e != nil {
			panic(e)
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

	handler := httpTransport.MakeHTTPHandler(instrumentingMiddleware, log.With(logger, "component", "HTTP"))

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	logger.Log("exit", <-errs)
}
