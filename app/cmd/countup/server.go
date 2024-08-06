package main

import (
	"context"
	"fmt"
	"time"

	apiv1 "github.com/jace-ys/countup/api/v1"
	genapi "github.com/jace-ys/countup/api/v1/gen/api"
	apipb "github.com/jace-ys/countup/api/v1/gen/grpc/api/pb"
	grpcapi "github.com/jace-ys/countup/api/v1/gen/grpc/api/server"
	httpapi "github.com/jace-ys/countup/api/v1/gen/http/api/server"
	httpweb "github.com/jace-ys/countup/api/v1/gen/http/web/server"
	genweb "github.com/jace-ys/countup/api/v1/gen/web"
	"github.com/jace-ys/countup/internal/app"
	"github.com/jace-ys/countup/internal/endpoints"
	"github.com/jace-ys/countup/internal/handler/api"
	"github.com/jace-ys/countup/internal/handler/web"
	"github.com/jace-ys/countup/internal/instrument"
	"github.com/jace-ys/countup/internal/postgres"
	"github.com/jace-ys/countup/internal/service/counter"
	counterstore "github.com/jace-ys/countup/internal/service/counter/store"
	"github.com/jace-ys/countup/internal/slog"
	"github.com/jace-ys/countup/internal/transport"
	"github.com/jace-ys/countup/internal/worker"
)

type ServerCmd struct {
	Port      int `env:"PORT" default:"8080" help:"Port for application server to listen on."`
	AdminPort int `env:"ADMIN_PORT" default:"9090" help:"Port for admin server to listen on."`

	OTLP struct {
		MetricsEndpoint string `env:"METRICS_ENDPOINT" default:"127.0.0.1:4317" help:"OTLP gRPC endpoint to send OpenTelemetry metrics to."`
		TracesEndpoint  string `env:"TRACES_ENDPOINT" default:"127.0.0.1:4317" help:"OTLP gRPC endpoint to send OpenTelemetry traces to."`
	} `embed:"" envprefix:"OTLP_" prefix:"otlp."`

	Database struct {
		ConnectionURI string `env:"CONNECTION_URI" required:"" help:"Connection URI for connecting to the database."`
	} `embed:"" envprefix:"DATABASE_" prefix:"database."`

	Worker struct {
		Concurrency int `env:"CONCURRENCY" default:"50" help:"Number of workers to run in the worker pool."`
	} `embed:"" envprefix:"WORKER_" prefix:"worker."`

	Counter struct {
		FinalizeWindow time.Duration `env:"FINALIZE_WINDOW" default:"1m" help:"Time period to wait before finalizing counter increments."`
	} `embed:"" envprefix:"COUNTER_" prefix:"counter."`
}

func (c *ServerCmd) Run(ctx context.Context, g *Globals) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	instrument.MustInitOTelProvider(ctx, genapi.APIName, genapi.APIVersion, c.OTLP.MetricsEndpoint, c.OTLP.TracesEndpoint)
	defer func() {
		if err := instrument.OTel.Shutdown(context.Background()); err != nil {
			slog.Error(ctx, "error shutting down otel provider", err)
		}
	}()

	db, err := postgres.NewPool(ctx, c.Database.ConnectionURI)
	if err != nil {
		return fmt.Errorf("init db pool: %w", err)
	}
	defer db.Close()

	worker, err := worker.NewPool(ctx, "app.worker", db, g.Writer, c.Worker.Concurrency)
	if err != nil {
		return fmt.Errorf("init worker pool: %w", err)
	}

	httpsrv := app.NewHTTPServer(ctx, "app.http", c.Port)
	grpcsrv := app.NewGRPCServer[apipb.APIServer](ctx, "app.grpc", c.Port+1)

	admin := app.NewAdminServer(ctx, c.AdminPort, g.Debug)
	admin.Administer(httpsrv, grpcsrv, worker)

	{
		countersvc := counter.New(db, worker, counterstore.New(), c.Counter.FinalizeWindow)
		// usersvc := counter.New(db, counterstore.New())

		handler, err := api.NewHandler(worker, countersvc)
		if err != nil {
			return fmt.Errorf("init handler: %w", err)
		}
		admin.Administer(handler)

		ep := endpoints.Goa(genapi.NewEndpoints).Adapt(handler)

		{
			transport := transport.GoaHTTP(httpapi.New, httpapi.Mount)
			httpsrv.RegisterHandler("/api/v1", transport.Adapt(ctx, ep, apiv1.OpenAPIFS))
		}
		{
			transport := transport.GoaGRPC(grpcapi.New)
			grpcsrv.RegisterHandler(&apipb.API_ServiceDesc, transport.Adapt(ctx, ep))
		}
	}

	{
		handler, err := web.NewHandler()
		if err != nil {
			return fmt.Errorf("init web: %w", err)
		}
		admin.Administer(handler)

		ep := endpoints.Goa(genweb.NewEndpoints).Adapt(handler)

		transport := transport.GoaHTTP(httpweb.New, httpweb.Mount)
		httpsrv.RegisterHandler("/", transport.Adapt(ctx, ep, web.StaticFS))
	}

	err = app.New(httpsrv, grpcsrv, admin, worker).Run(ctx)
	if err != nil {
		slog.Error(ctx, "encountered error while running app", err)
		return fmt.Errorf("app run: %w", err)
	}

	return nil
}
