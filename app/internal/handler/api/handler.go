package api

import (
	"context"

	"github.com/alexliesenfeld/health"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/healthz"
	"github.com/jace-ys/countup/internal/service/counter"
	"github.com/jace-ys/countup/internal/worker"
)

var _ api.Service = (*Handler)(nil)

type Handler struct {
	workers *worker.Pool
	counter CounterService
}

type CounterService interface {
	GetInfo(ctx context.Context) (*counter.Info, error)
	RequestIncrement(ctx context.Context, user string) error
}

func NewHandler(workers *worker.Pool, counter CounterService) (*Handler, error) {
	worker.Register(workers, &EchoWorker{})

	return &Handler{
		workers: workers,
		counter: counter,
	}, nil
}

var _ healthz.Target = (*Handler)(nil)

func (h *Handler) HealthChecks() []health.Check {
	return []health.Check{
		{
			Name: "handler:countup",
			Check: func(ctx context.Context) error {
				return nil
			},
		},
	}
}
