package slog

import (
	"context"
	"errors"
	"log/slog"

	"goa.design/clue/log"
)

type StdSlogHandler struct {
	ctx   context.Context
	level slog.Level
	group string
}

func AsStdSlogHandler(ctx context.Context, level slog.Level) *StdSlogHandler {
	return &StdSlogHandler{
		ctx:   log.Context(ctx, log.WithDebug()),
		level: level,
	}
}

var _ slog.Handler = (*StdSlogHandler)(nil)

func (l *StdSlogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= l.level
}

func (l *StdSlogHandler) Handle(ctx context.Context, record slog.Record) error {
	var attrs []log.Fielder
	record.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, KV(a.Key, a.Value.Any()))
		return true
	})

	logCtx := log.WithContext(ctx, l.ctx)
	logCtx = log.With(logCtx, attrs...)

	switch record.Level {
	case slog.LevelDebug:
		Debug(logCtx, record.Message)
	case slog.LevelInfo:
		Info(logCtx, record.Message)
	case slog.LevelWarn:
		Error(logCtx, "warn", errors.New(record.Message))
	case slog.LevelError:
		Error(logCtx, "error", errors.New(record.Message))
	default:
		Print(logCtx, record.Message)
	}
	return nil
}

func (l *StdSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	kvs := make([]log.Fielder, len(attrs))
	for i, attr := range attrs {
		if l.group != "" {
			k := l.group + "." + attr.Key
			kvs[i] = KV(k, attr.Value.Any())
		} else {
			kvs[i] = KV(attr.Key, attr.Value.Any())
		}
	}

	cp := l.clone()
	cp.ctx = log.With(cp.ctx, kvs...)
	return cp
}

func (l *StdSlogHandler) WithGroup(name string) slog.Handler {
	cp := l.clone()
	cp.group = name
	return cp
}

func (l *StdSlogHandler) clone() *StdSlogHandler {
	cp := *l
	return &cp
}

type NopHandler struct {
}

func AsNopHandler() *NopHandler {
	return &NopHandler{}
}

var _ slog.Handler = (*NopHandler)(nil)

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
