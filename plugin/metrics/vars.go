package vars

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Request* and Response* are the prometheus counters and gauges we are using for exporting metrics.
var (
	RequestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: subsystem,
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "Counter of DNS requests made per zone, protocol and family.",
	}, []string{"server", "zone", "proto", "family", "type"})

	// Top10 client IPs last 5 minutes.
	ClientCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		// Namespace: plugin.Namespace,
		Namespace: subsystem,
		Name:      "client_count_recent",
		Help:      "Number of times client accessed recently",
	}, []string{"client"})
)

const (
	subsystem = "waf"

	// Dropped indicates we dropped the query before any handling. It has no closing dot, so it can not be a valid zone.
	Dropped = "dropped"
)
