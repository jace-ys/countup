package api

import (
	"context"

	"github.com/alexliesenfeld/health"
	"github.com/markbates/goth"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/healthz"
	"github.com/jace-ys/countup/internal/service/counter"
)

var _ api.Service = (*Handler)(nil)

type Handler struct {
	authn   goth.Provider
	counter CounterService
}

type CounterService interface {
	GetInfo(ctx context.Context) (*counter.Info, error)
	RequestIncrement(ctx context.Context, user string) error
}

func NewHandler(authn goth.Provider, counter CounterService) (*Handler, error) {
	return &Handler{
		authn:   authn,
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
