package ctxlog_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jace-ys/countup/internal/ctxlog"
)

func TestSlogHandler(t *testing.T) {
	buf := &bytes.Buffer{}
	ctx := ctxlog.NewContext(t.Context(), buf, false)

	logger := slog.New(ctxlog.AsSlogHandler(ctx, slog.LevelDebug))

	logger.Debug("test debug", "count", 1.0)
	logger.Info("test info", "count", 2.0)
	logger.Warn("test warn", "count", 3.0)
	logger.Error("test error", "count", 4.0)

	assertLogOutput(t, buf, []map[string]any{
		{"level": "debug", "msg": "test debug", "count": 1.0},
		{"level": "info", "msg": "test info", "count": 2.0},
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0},
	})
}

func TestSlogHandlerLevel(t *testing.T) {
	buf := &bytes.Buffer{}
	ctx := ctxlog.NewContext(t.Context(), buf, false)

	logger := slog.New(ctxlog.AsSlogHandler(ctx, slog.LevelWarn))

	logger.DebugContext(ctx, "test debug", "count", 1.0)
	logger.InfoContext(ctx, "test info", "count", 2.0)
	logger.WarnContext(ctx, "test warn", "count", 3.0)
	logger.ErrorContext(ctx, "test error", "count", 4.0)

	assertLogOutput(t, buf, []map[string]any{
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0},
	})
}

func TestSlogHandlerContext(t *testing.T) {
	buf := &bytes.Buffer{}
	ctx := ctxlog.NewContext(t.Context(), buf, false)
	ctx = ctxlog.With(ctx, ctxlog.KV("foo", "bar"))

	logger := slog.New(ctxlog.AsSlogHandler(ctx, slog.LevelDebug))

	logger.DebugContext(ctx, "test debug", "count", 1.0)
	logger.InfoContext(ctx, "test info", "count", 2.0)
	logger.WarnContext(ctx, "test warn", "count", 3.0)
	logger.ErrorContext(ctx, "test error", "count", 4.0)

	ctx = ctxlog.With(ctx, ctxlog.KV("ping", "pong"))

	logger.DebugContext(ctx, "test debug", "count", 1.0)
	logger.InfoContext(ctx, "test info", "count", 2.0)
	logger.WarnContext(ctx, "test warn", "count", 3.0)
	logger.ErrorContext(ctx, "test error", "count", 4.0)

	assertLogOutput(t, buf, []map[string]any{
		{"level": "debug", "msg": "test debug", "count": 1.0, "foo": "bar"},
		{"level": "info", "msg": "test info", "count": 2.0, "foo": "bar"},
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0, "foo": "bar"},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0, "foo": "bar"},
		{"level": "debug", "msg": "test debug", "count": 1.0, "foo": "bar"},
		{"level": "info", "msg": "test info", "count": 2.0, "foo": "bar"},
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0, "foo": "bar"},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0, "foo": "bar"},
	})
}

func TestSlogHandlerWith(t *testing.T) {
	buf := &bytes.Buffer{}
	ctx := ctxlog.NewContext(t.Context(), buf, false)

	logger := slog.New(ctxlog.AsSlogHandler(ctx, slog.LevelDebug))

	logger.DebugContext(ctx, "test debug", "count", 1.0)
	logger.InfoContext(ctx, "test info", "count", 2.0)
	logger.WarnContext(ctx, "test warn", "count", 3.0)
	logger.ErrorContext(ctx, "test error", "count", 4.0)

	logger = logger.With("ping", "pong")

	logger.DebugContext(ctx, "test debug", "count", 1.0)
	logger.InfoContext(ctx, "test info", "count", 2.0)
	logger.WarnContext(ctx, "test warn", "count", 3.0)
	logger.ErrorContext(ctx, "test error", "count", 4.0)

	assertLogOutput(t, buf, []map[string]any{
		{"level": "debug", "msg": "test debug", "count": 1.0},
		{"level": "info", "msg": "test info", "count": 2.0},
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0},
		{"level": "debug", "msg": "test debug", "count": 1.0, "ping": "pong"},
		{"level": "info", "msg": "test info", "count": 2.0, "ping": "pong"},
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0, "ping": "pong"},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0, "ping": "pong"},
	})
}

func TestSlogHandlerGroup(t *testing.T) {
	buf := &bytes.Buffer{}
	ctx := ctxlog.NewContext(t.Context(), buf, false)

	logger := slog.New(ctxlog.AsSlogHandler(ctx, slog.LevelDebug))

	logger.DebugContext(ctx, "test debug", "count", 1.0)
	logger.InfoContext(ctx, "test info", "count", 2.0)
	logger.WarnContext(ctx, "test warn", "count", 3.0)
	logger.ErrorContext(ctx, "test error", "count", 4.0)

	logger = logger.WithGroup("foo").With("bar", "baz").WithGroup("ping").With("pong", "table")

	logger.DebugContext(ctx, "test debug", "count", 1.0)
	logger.InfoContext(ctx, "test info", "count", 2.0)
	logger.WarnContext(ctx, "test warn", "count", 3.0)
	logger.ErrorContext(ctx, "test error", "count", 4.0)

	assertLogOutput(t, buf, []map[string]any{
		{"level": "debug", "msg": "test debug", "count": 1.0},
		{"level": "info", "msg": "test info", "count": 2.0},
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0},
		{"level": "debug", "msg": "test debug", "count": 1.0, "foo.bar": "baz", "ping.pong": "table"},
		{"level": "info", "msg": "test info", "count": 2.0, "foo.bar": "baz", "ping.pong": "table"},
		{"level": "error", "msg": "warn", "err": "test warn", "count": 3.0, "foo.bar": "baz", "ping.pong": "table"},
		{"level": "error", "msg": "error", "err": "test error", "count": 4.0, "foo.bar": "baz", "ping.pong": "table"},
	})
}

func TestSlogHandlerConcurrent(t *testing.T) {
	ctx := ctxlog.NewContext(t.Context(), io.Discard, false)
	logger := slog.New(ctxlog.AsSlogHandler(ctx, slog.LevelDebug))

	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if id%2 == 0 {
				logger.Info("table", "ping", "pong")
			} else {
				logger.WithGroup("foo").With("bar", "baz").Info("table", "ping", "pong")
			}
		}(i)
	}

	wg.Wait()
}

func assertLogOutput(t *testing.T, r io.Reader, expected []map[string]any) {
	scanner := bufio.NewScanner(r)
	var actual []map[string]any

	for scanner.Scan() {
		var data map[string]any
		require.NoError(t, json.Unmarshal(scanner.Bytes(), &data))
		delete(data, "time")
		delete(data, "file")
		actual = append(actual, data)
	}

	require.Len(t, actual, len(expected))

	for i, data := range actual {
		assert.Equal(t, expected[i], data)
	}
}
