package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"time"
)

type Prometheus struct {
	requestCounter *prometheus.CounterVec
	requestLatency *prometheus.HistogramVec
}

func (p *Prometheus) HandleFunc() gin.HandlerFunc {
	return func (c *gin.Context) {
		start := time.Now()
		c.Next()

		end := time.Since(start)
		if c.Request.URL.String() != "/metrics" {
			status := strconv.Itoa(c.Writer.Status())
			p.requestCounter.WithLabelValues(c.Request.Method, status, c.Request.URL.String()).Inc()
			p.requestLatency.WithLabelValues(c.Request.Method, status, c.Request.URL.String()).Observe(end.Seconds())
		}
	}
}

func NewPrometheus(namespace, context string) *Prometheus {
	p := &Prometheus{}
	c := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: context,
			Name: "requests_total",
			Help: "the number of HTTP requests processed",
		},
		[]string{"method", "status", "path"},
	)

	if err := prometheus.Register(c); err != nil {
		log.Printf("could register counter: %s\n", err)
	}

	h := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: context,
			Name: "requests_latency_seconds",
			Help: "latency of HTTP requests processed",
		},
		[]string{"method", "status", "path"},
	)

	if err := prometheus.Register(h); err != nil {
		log.Printf("could register histogram: %s\n", err)
	}

	p.requestCounter = c
	p.requestLatency = h

	return p
}