package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestCounter     *prometheus.CounterVec
	HttpRequestLatency     *prometheus.HistogramVec
	HttpRequestConcurrency *prometheus.GaugeVec
)

func init() {
	HttpRequestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_counter",
		Help: "each http request counter",
	}, []string{"path", "method", "status"})

	HttpRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_latency_sec",
		Help:    "http history request duration distribution",
		Buckets: []float64{0.05, 0.2, 0.5, 1, 5, 10, 30},
	}, []string{"path", "method"})

	HttpRequestConcurrency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_request_in_flight",
		Help: "http request concurrency number",
	}, []string{"path", "method"})
}
