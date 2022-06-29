package state

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const MaxTxExecPeriod int64 = 2000000

// Metrics represents the txpool metrics
type Metrics struct {
	// Pending transactions
	TxnExceedPeriod metrics.Counter
}

// GetPrometheusMetrics return the txpool metrics instance
func GetPrometheusMetrics(namespace string, labelsWithValues ...string) *Metrics {
	labels := []string{}

	for i := 0; i < len(labelsWithValues); i += 2 {
		labels = append(labels, labelsWithValues[i])
	}

	return &Metrics{
		TxnExceedPeriod: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "state",
			Name:      "txn_exceed_period",
			Help:      "Number of transactions whose execution period exceeds MaxTxExecPeriod.",
		}, labels).With(labelsWithValues...),
	}
}

// NilMetrics will return the non operational txpool metrics
func NilMetrics() *Metrics {
	return &Metrics{
		TxnExceedPeriod: discard.NewCounter(),
	}
}
