// DON'T EDIT THIS FILE!

package newapp

const tmplMetricMetric = `// Package metric contains metrics for the app
package metric

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"

	"{{ .ModuleName }}/config"
)

// Metric define the use case interface
type Metric interface {
	SaveCLI(name string, duration float64) error
	SaveHTTP(*http)
}

// Middleware define the middleware metrics
func Middleware(m Metric) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app := newHTTP(ctx.Request.URL.String(), ctx.Request.Method)
		app.started()
		ctx.Next()
		app.finished()
		app.statusCode = strconv.Itoa(ctx.Writer.Status())
		m.SaveHTTP(app)
	}
}

type metric struct {
	cfg *config.Config

	cmdHistogram         *prometheus.HistogramVec
	httpRequestHistogram *prometheus.HistogramVec
}

// NewPrometheusMetrics create a new prometheus service
func NewPrometheusMetrics(cfg *config.Config) (*metric, error) {
	m := &metric{
		cfg: cfg,
		cmdHistogram: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "pushgateway",
			Name:      "cmd_duration_seconds",
			Help:      "CLI application execution in seconds",
			Buckets:   prometheus.DefBuckets,
		}, []string{"name"}),

		httpRequestHistogram: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "http",
			Name:      "request_duration_seconds",
			Help:      "The latency of the HTTP requests.",
			Buckets:   prometheus.DefBuckets,
		}, []string{"handler", "method", "code"}),
	}

	var msg = "duplicate metrics collector registration attempted"
	if err := prometheus.Register(m.cmdHistogram); err != nil && err.Error() != msg {
		return nil, err
	}

	if err := prometheus.Register(m.httpRequestHistogram); err != nil && err.Error() != msg {
		return nil, err
	}

	return m, nil
}

type cli struct {
	name       string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

// NewCLI create a new cli
func NewCLI(name string) *cli {
	return &cli{name: name}
}

// Started set the started time
func (c *cli) Started() {
	c.StartedAt = time.Now()
}

// Finished set the finished time
func (c *cli) Finished() {
	c.FinishedAt = time.Now()
	c.Duration = time.Since(c.StartedAt).Seconds()
}

// SaveCLI send metrics to server
func (s *metric) SaveCLI(name string, duration float64) error {
	s.cmdHistogram.WithLabelValues(name).Observe(duration)
	return push.New(s.cfg.GetApplication().GetPrometheusPushgateway(), "cmd_job").Collector(s.cmdHistogram).Push()
}

type http struct {
	handler    string
	method     string
	statusCode string
	startedAt  time.Time
	finishedAt time.Time
	duration   float64
}

// SaveHTTP send metrics to server
func (s *metric) SaveHTTP(h *http) {
	s.httpRequestHistogram.WithLabelValues(h.handler, h.method, h.statusCode).Observe(h.duration)
}

func newHTTP(handler string, method string) *http {
	return &http{handler: handler, method: method}
}

func (h *http) started() {
	h.startedAt = time.Now()
}

func (h *http) finished() {
	h.finishedAt = time.Now()
	h.duration = time.Since(h.startedAt).Seconds()
}

`
