package slog

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"goa.design/clue/log"
)

var _ slog.Handler = (*StdSlogHandler)(nil)

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

func (l *StdSlogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= l.level
}

func (l *StdSlogHandler) Handle(ctx context.Context, record slog.Record) error {
	var attrs []log.Fielder
	record.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, KV(a.Key, a.Value.Any()))
		return true
	})

	ctx = log.WithContext(ctx, l.ctx)
	ctx = log.With(ctx, attrs...)

	switch record.Level {
	case slog.LevelDebug:
		Debug(ctx, record.Message)
	case slog.LevelInfo:
		Info(ctx, record.Message)
	case slog.LevelWarn:
		Error(ctx, "warn", errors.New(record.Message))
	case slog.LevelError:
		Error(ctx, "error", errors.New(record.Message))
	default:
		Print(ctx, record.Message)
	}
	return nil
}

func (l *StdSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	kvs := make([]log.Fielder, len(attrs))
	for i, attr := range attrs {
		if l.group != "" {
			kvs[i] = KV(fmt.Sprintf("%s.%s", l.group, attr.Key), attr.Value.Any())
		} else {
			kvs[i] = KV(attr.Key, attr.Value.Any())
		}
	}

	cp := l.clone()
	cp.ctx = log.With(l.ctx, kvs...)
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
