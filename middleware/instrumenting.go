package middleware

import (
	"fmt"
	"github.com/pashukhin/coins-test-task/entity"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"time"

	"github.com/go-kit/kit/metrics"
)

func NewInstrumentingMiddleware() Middleware {
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

	return &instrumentingMiddleware{
		&middleware{},
		requestCount,
		requestLatency,
	}
}

type instrumentingMiddleware struct {
	*middleware
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

func (mw *instrumentingMiddleware) Accounts() (output []*entity.Account, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "accounts", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Accounts()
	return
}

func (mw *instrumentingMiddleware) Payments() (output []*entity.Payment, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "payments", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Payments()
	return
}

func (mw *instrumentingMiddleware) Send(fromID, toID int64, amount float64) (output *entity.Payment, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "send", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Send(fromID, toID, amount)
	return
}

func (mw *instrumentingMiddleware) Account(id int64) (output *entity.Account, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "account", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Account(id)
	return
}
