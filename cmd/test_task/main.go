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

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {

	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address")
	dbHost := flag.String("db.host", "db", "db host")
	dbPort := flag.Int("db.port", 5432, "db port")
	dbUser := flag.String("db.user", "user", "db user")
	dbPassword := flag.String("db.password", "password", "db password")
	dbDatabase := flag.String("db.database", "db", "database name")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)

	//db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5432 user=user password=password dbname=db sslmode=disable")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *dbHost, *dbPort, *dbUser, *dbPassword, *dbDatabase)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
    driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file://db/migrations",
        "postgres", driver)
    if err != nil {
		panic(err)
	}
    if err := m.Up(); err != nil {
    	if err != migrate.ErrNoChange {
			panic(err)
		}
	}

	accounts := repository.NewAccountRepository(db)
	payments := repository.NewPaymentRepository(db)
	service := business.NewService()
	service.SetAccountRepository(accounts)
	service.SetPaymentRepository(payments)

	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	loggingMiddleware.SetNext(service)

	handler := httpTransport.MakeHTTPHandler(loggingMiddleware, log.With(logger, "component", "HTTP"))

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
