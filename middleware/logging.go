package middleware

import (
	"github.com/pashukhin/coins-test-task/entity"
	"time"

	"github.com/go-kit/kit/log"
)

func NewLoggingMiddleware(logger log.Logger) Middleware {
	return &loggingMiddleware{middleware: &middleware{}, logger: logger}
}

type loggingMiddleware struct {
	*middleware
	logger log.Logger
}

func (mw *loggingMiddleware) Accounts() (output []*entity.Account, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "accounts",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Accounts()
	return
}

func (mw *loggingMiddleware) Payments() (output []*entity.Payment, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "payments",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Payments()
	return
}

func (mw *loggingMiddleware) Send(fromID, toID int64, amount float64) (output *entity.Payment, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "send",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Send(fromID, toID, amount)
	return
}

func (mw *loggingMiddleware) Account(id int64) (output *entity.Account, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "account",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Account(id)
	return
}
