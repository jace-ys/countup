package api

import (
	"context"

	"github.com/alexliesenfeld/health"
	"github.com/markbates/goth"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/healthz"
	"github.com/jace-ys/countup/internal/service/counter"
	"github.com/jace-ys/countup/internal/service/user"
)

var _ api.Service = (*Handler)(nil)

type Handler struct {
	authn   goth.Provider
	counter CounterService
	users   UserService

	jwtSigningSecret []byte
}

type CounterService interface {
	GetInfo(ctx context.Context) (*counter.Info, error)
	RequestIncrement(ctx context.Context, user string) error
}

type UserService interface {
	GetUser(ctx context.Context, id string) (*user.User, error)
	CreateUserIfNotExists(ctx context.Context, email string) (*user.User, error)
}

func NewHandler(authn goth.Provider, counter CounterService, users UserService) (*Handler, error) {
	return &Handler{
		authn:            authn,
		counter:          counter,
		users:            users,
		jwtSigningSecret: []byte("secret"),
	}, nil
}

var _ healthz.Target = (*Handler)(nil)

func (h *Handler) HealthChecks() []health.Check {
	return []health.Check{
		{
			Name: "handler:api",
			Check: func(ctx context.Context) error {
				return nil
			},
		},
	}
}
