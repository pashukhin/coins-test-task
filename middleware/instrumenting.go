package main

import (
	"fmt"
	"github.com/pashukhin/coins-test-task/entity"
	"github.com/pashukhin/coins-test-task/service"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           service.Service
}

func (mw instrumentingMiddleware) Accounts() (output []*entity.Account, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "accounts", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Accounts()
	return
}

func (mw instrumentingMiddleware) Payments() (output []*entity.Payment, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "payments", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Payments()
	return
}

func (mw instrumentingMiddleware) Send(fromID, toID int64, amount float64) (output *entity.Payment, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "send", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Send(fromID, toID, amount)
	return
}
