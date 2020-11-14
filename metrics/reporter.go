// Package metrics is for instrumentation and debugging
package metrics

import "time"

// Tags is a map of fields to add to a metric:
type Tags map[string]string

var (
	DefaultReporter Reporter = NewReporter()
)

// Reporter is an interface for collecting and instrumenting metrics
type Reporter interface {
	Init(...Option) error
	Count(id string, value int64, tags Tags) error
	Gauge(id string, value float64, tags Tags) error
	Timing(id string, value time.Duration, tags Tags) error
	Options() Options
}
