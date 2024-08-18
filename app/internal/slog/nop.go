package slog

import (
	"context"
	"log/slog"
)

type NopHandler struct {
}

func AsNopHandler(ctx context.Context) *NopHandler {
	return &NopHandler{}
}

func (l NopHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return false
}

func (l NopHandler) Handle(ctx context.Context, record slog.Record) error {
	return nil
}

func (l NopHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return l
}

func (l NopHandler) WithGroup(name string) slog.Handler {
	return l
}
