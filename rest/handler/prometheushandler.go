package handler

import (
	"net/http"
	"strconv"

	"github.com/gzorm/commons/core/metric"
	"github.com/gzorm/commons/core/timex"
	"github.com/gzorm/commons/rest/internal/response"
)

const serverNamespace = "http_server"

var (
	metricServerReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "http server requests duration(ms).",
		Labels:    []string{"path", "method", "code"},
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 750, 1000},
	})

	metricServerReqCodeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "http server requests error count.",
		Labels:    []string{"path", "method", "code"},
	})
)

// PrometheusHandler returns a middleware that reports stats to prometheus.
func PrometheusHandler(path, method string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := timex.Now()
			cw := response.NewWithCodeResponseWriter(w)
			defer func() {
				code := strconv.Itoa(cw.Code)
				metricServerReqDur.Observe(timex.Since(startTime).Milliseconds(), path, method, code)
				metricServerReqCodeTotal.Inc(path, method, code)
			}()

			next.ServeHTTP(cw, r)
		})
	}
}
