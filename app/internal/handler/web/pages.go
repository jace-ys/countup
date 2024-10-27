package web

import (
	"bytes"
	"context"
	"fmt"
)

func (h *Handler) Index(ctx context.Context) ([]byte, error) {
	data := struct {
		Name string
	}{
		Name: "Count Up!",
	}

	return h.render("index.html", data)
}

func (h *Handler) Another(ctx context.Context) ([]byte, error) {
	data := struct {
		Name string
	}{
		Name: "Page",
	}

	return h.render("another.html", data)
}

type renderData struct {
	Data any
}

func (h *Handler) render(page string, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := h.tmpls.ExecuteTemplate(buf, page, renderData{data}); err != nil {
		return nil, fmt.Errorf("render: %w", err)
	}
	return buf.Bytes(), nil
}
