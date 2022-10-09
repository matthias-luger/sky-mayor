package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	electionPeriodInsertCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sky_mayor_election_period_inserted",
		Help: "Count of inserted election periods",
	})

	votingInsertCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sky_mayor_votings_inserted",
		Help: "Count of inserted voting data",
	})

	errorCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sky_mayor_errors",
		Help: "Count of errors",
	})
)

func Init() error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":2112", nil)
}

func VotingInserted() {
	votingInsertCounter.Inc()
}

func ElectionPeriodInserted() {
	electionPeriodInsertCounter.Inc()
}

func AddError() {
	errorCounter.Inc()
}
