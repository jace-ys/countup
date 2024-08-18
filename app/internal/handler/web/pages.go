package web

import (
	"context"
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
