package web

import (
	"context"
	"embed"
	"fmt"
	"html/template"

	"github.com/alexliesenfeld/health"
	"github.com/gorilla/securecookie"
	"github.com/markbates/goth"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/api/v1/gen/web"
	"github.com/jace-ys/countup/internal/healthz"
)

var (
	//go:embed static/*
	StaticFS embed.FS

	//go:embed templates/*
	templateFS embed.FS
)

var _ web.Service = (*Handler)(nil)

type Handler struct {
	authn   goth.Provider
	cookies *securecookie.SecureCookie
	api     *api.Client
	tmpls   *template.Template
}

func NewHandler(authn goth.Provider, cookies *securecookie.SecureCookie, api *api.Client) (*Handler, error) {
	tmpls, err := template.New("tmpls").ParseFS(templateFS, "**/*.html")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	return &Handler{
		authn:   authn,
		cookies: cookies.MaxAge(86400),
		api:     api,
		tmpls:   tmpls,
	}, nil
}

var _ healthz.Target = (*Handler)(nil)

func (h *Handler) HealthChecks() []health.Check {
	return []health.Check{
		{
			Name: "handler:web",
			Check: func(ctx context.Context) error {
				return nil
			},
		},
	}
}
