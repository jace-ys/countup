package healthz

import (
	"time"

	"github.com/alexliesenfeld/health"
)

type Target interface {
	HealthChecks() []health.Check
}

func NewChecker(checks ...health.Check) health.Checker {
	opts := []health.CheckerOption{
		health.WithDisabledAutostart(),
	}

	for _, check := range checks {
		check.Timeout = time.Second
		opts = append(opts, health.WithCheck(check))
	}

	return health.NewChecker(opts...)
}