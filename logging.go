package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) Accounts() (output []*Account, err error) {
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

func (mw loggingMiddleware) Payments() (output []*Payment, err error) {
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

func (mw loggingMiddleware) Send(fromID, toID string, amount float64) (output *Payment, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "send",
			"parameters", fromID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Send(fromID, toID, amount)
	return
}
