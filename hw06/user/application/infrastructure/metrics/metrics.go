package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"time"
)

type Prometheus struct {
	requestCounter 	*prometheus.CounterVec
	requestLatency  *prometheus.HistogramVec

	MetricsPath string
}

func (p *Prometheus) HandleFunc() gin.HandlerFunc {
	return func (c *gin.Context) {
		start := time.Now()
		c.Next()

		end := time.Since(start)
		if p.MetricsPath != c.Request.URL.String() {
			status := strconv.Itoa(c.Writer.Status())
			p.requestCounter.WithLabelValues(c.Request.Method, status, c.Request.URL.String()).Inc()
			p.requestLatency.WithLabelValues(c.Request.Method, status, c.Request.URL.String()).Observe(end.Seconds())
		}
	}
}

func NewPrometheus(namespsace, context, path string) *Prometheus {
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespsace,
			Subsystem: context,
			Name: "requests_total",
			Help: "the number of HTTP requests processed",
		},
		[]string{"method", "status", "path"},
	)

	requestLatency := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespsace,
			Subsystem: context,
			Name: "requests_latency_seconds",
			Help: "latency of HTTP requests processed",
		},
		[]string{"method", "status", "path"},
	)

	if err := prometheus.Register(requestCounter); err != nil {
		log.Printf("could not be registered: %s", err)
	}

	if err := prometheus.Register(requestLatency); err != nil {
		log.Printf("could not be registered: %s", err)
	}

	p := &Prometheus{
		requestCounter: requestCounter,
		requestLatency: requestLatency,
		MetricsPath: path,
	}

	return p
}