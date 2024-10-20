package teapot

import (
	"context"

	"github.com/alexliesenfeld/health"

	"github.com/jace-ys/countup/api/v1/gen/teapot"
	"github.com/jace-ys/countup/internal/healthz"
	"github.com/jace-ys/countup/internal/worker"
)

var _ teapot.Service = (*Handler)(nil)

type Handler struct {
	workers *worker.Pool
}

func NewHandler(workers *worker.Pool) (*Handler, error) {
	worker.Register(workers, &EchoWorker{})

	return &Handler{
		workers: workers,
	}, nil
}

var _ healthz.Target = (*Handler)(nil)

func (h *Handler) HealthChecks() []health.Check {
	return []health.Check{
		{
			Name: "handler:teapot",
			Check: func(ctx context.Context) error {
				return nil
			},
		},
	}
}
