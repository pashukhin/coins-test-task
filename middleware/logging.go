package main

import (
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/service"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func (mw loggingMiddleware) Accounts() (output []*entity.Account, err error) {
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

func (mw loggingMiddleware) Payments() (output []*entity.Payment, err error) {
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

func (mw loggingMiddleware) Send(fromID, toID int64, amount float64) (output *entity.Payment, err error) {
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
