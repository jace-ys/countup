package web

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"html/template"

	"github.com/alexliesenfeld/health"

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
	tmpls *template.Template
}

func NewHandler() (*Handler, error) {
	tmpls, err := template.New("tmpls").ParseFS(templateFS, "**/*.html")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	return &Handler{
		tmpls: tmpls,
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

var commonVars = struct {
	BaseAPIEndpoint string
}{
	BaseAPIEndpoint: "/api/v1",
}

type renderData struct {
	Vars any
	Data any
}

func (h *Handler) render(page string, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := h.tmpls.ExecuteTemplate(buf, page, renderData{commonVars, data})
	return buf.Bytes(), err
}
