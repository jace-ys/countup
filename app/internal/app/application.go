package app

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/jace-ys/countup/internal/slog"
)

type Application struct {
	servers []Server
}

func New(servers ...Server) *Application {
	return &Application{
		servers: servers,
	}
}

type Server interface {
	Name() string
	Kind() string
	Addr() string
	Serve(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

func (a *Application) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, srv := range a.servers {
		g.Go(func() error {
			slog.Print(ctx, "server listening",
				slog.KV("server", srv.Name()),
				slog.KV("kind", srv.Kind()),
				slog.KV("addr", srv.Addr()),
			)
			return srv.Serve(ctx)
		})
	}

	slog.Print(ctx, "application started")
	<-ctx.Done()
	slog.Print(ctx, "application shutting down gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, srv := range a.servers {
		g.Go(func() error {
			if err := srv.Shutdown(shutdownCtx); err != nil {
				slog.Error(ctx, "server shutdown error", err,
					slog.KV("server", srv.Name()),
					slog.KV("kind", srv.Kind()),
				)
			}
			slog.Print(ctx, "server shutdown complete",
				slog.KV("server", srv.Name()),
				slog.KV("kind", srv.Kind()),
			)
			return nil
		})
	}

	defer slog.Print(ctx, "application stopped")
	return g.Wait() //nolint:wrapcheck
}
