package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/markbates/goth/providers/google"

	apiv1 "github.com/jace-ys/countup/api/v1"
	genapi "github.com/jace-ys/countup/api/v1/gen/api"
	grpcapiclient "github.com/jace-ys/countup/api/v1/gen/grpc/api/client"
	apipb "github.com/jace-ys/countup/api/v1/gen/grpc/api/pb"
	grpcapi "github.com/jace-ys/countup/api/v1/gen/grpc/api/server"
	teapotpb "github.com/jace-ys/countup/api/v1/gen/grpc/teapot/pb"
	grpcteapot "github.com/jace-ys/countup/api/v1/gen/grpc/teapot/server"
	httpapi "github.com/jace-ys/countup/api/v1/gen/http/api/server"
	httpteapot "github.com/jace-ys/countup/api/v1/gen/http/teapot/server"
	httpweb "github.com/jace-ys/countup/api/v1/gen/http/web/server"
	genteapot "github.com/jace-ys/countup/api/v1/gen/teapot"
	genweb "github.com/jace-ys/countup/api/v1/gen/web"
	"github.com/jace-ys/countup/internal/app"
	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/endpoint"
	"github.com/jace-ys/countup/internal/handler/api"
	"github.com/jace-ys/countup/internal/handler/teapot"
	"github.com/jace-ys/countup/internal/handler/web"
	"github.com/jace-ys/countup/internal/instrument"
	"github.com/jace-ys/countup/internal/postgres"
	"github.com/jace-ys/countup/internal/service/counter"
	counterstore "github.com/jace-ys/countup/internal/service/counter/store"
	"github.com/jace-ys/countup/internal/service/user"
	userstore "github.com/jace-ys/countup/internal/service/user/store"
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

	OAuth struct {
		ClientID     string `env:"CLIENT_ID" required:"" help:"Client ID for the Google OAuth2 configuration."`
		ClientSecret string `env:"CLIENT_SECRET" required:"" help:"Client secret for the Google OAuth2 configuration."`
		RedirectURL  string `env:"REDIRECT_URL" default:"http://localhost:8080/login/google/callback" help:"URL to redirect to upon successful OAuth2 authentication."`
	} `embed:"" envprefix:"OAUTH_" prefix:"oauth."`

	Counter struct {
		FinalizeWindow time.Duration `env:"FINALIZE_WINDOW" default:"1m" help:"Time period to wait before finalizing counter increments."`
	} `embed:"" envprefix:"COUNTER_" prefix:"counter."`
}

func (c *ServerCmd) Run(ctx context.Context, g *Globals) error {
	instrument.MustInitOTelProvider(ctx, genapi.APIName, genapi.APIVersion, c.OTLP.MetricsEndpoint, c.OTLP.TracesEndpoint)
	defer func() {
		if err := instrument.OTel.Shutdown(ctx); err != nil {
			ctxlog.Error(ctx, "error shutting down otel provider", err)
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

	authn := google.New(c.OAuth.ClientID, c.OAuth.ClientSecret, c.OAuth.RedirectURL, "https://www.googleapis.com/auth/userinfo.email")

	counterSvc := counter.New(db, worker, counterstore.New(), c.Counter.FinalizeWindow)
	userSvc := user.New(db, userstore.New())

	httpSrv := app.NewHTTPServer(ctx, "app.http", c.Port)
	grpcSrv := app.NewGRPCServer[apipb.APIServer](ctx, "app.grpc", c.Port+1)

	admin := app.NewAdminServer(ctx, c.AdminPort, g.Debug)
	admin.Administer(httpSrv, grpcSrv, worker)

	{
		handler, err := api.NewHandler(authn, counterSvc, userSvc)
		if err != nil {
			return fmt.Errorf("init api handler: %w", err)
		}
		admin.Administer(handler)

		ep := endpoint.Goa(genapi.NewEndpoints).Adapt(handler)

		{
			transport := transport.GoaHTTP(httpapi.New, httpapi.Mount)
			httpSrv.RegisterHandler("/api/v1", transport.Adapt(ctx, ep, apiv1.OpenAPIFS))
		}
		{
			transport := transport.GoaGRPC(grpcapi.New)
			grpcSrv.RegisterHandler(&apipb.API_ServiceDesc, transport.Adapt(ctx, ep))
		}
	}

	{
		handler, err := teapot.NewHandler(worker)
		if err != nil {
			return fmt.Errorf("init teapot handler: %w", err)
		}
		admin.Administer(handler)

		ep := endpoint.Goa(genteapot.NewEndpoints).Adapt(handler)

		{
			transport := transport.GoaHTTP(httpteapot.New, httpteapot.Mount)
			httpSrv.RegisterHandler("/teapot", transport.Adapt(ctx, ep, apiv1.OpenAPIFS))
		}
		{
			transport := transport.GoaGRPC(grpcteapot.New)
			grpcSrv.RegisterHandler(&teapotpb.Teapot_ServiceDesc, transport.Adapt(ctx, ep))
		}
	}

	{
		cookies := securecookie.New([]byte("secret-hash-key"), []byte("secret-block-key"))

		tc, err := transport.GoaGRPCClient(grpcapiclient.NewClient).Adapt(grpcSrv.Addr())
		if err != nil {
			return fmt.Errorf("init api client: %w", err)
		}
		defer tc.Close()

		apiClient := genapi.NewClient(tc.Client().AuthToken(), tc.Client().CounterGet(), tc.Client().CounterIncrement())

		handler, err := web.NewHandler(authn, cookies, apiClient)
		if err != nil {
			return fmt.Errorf("init web handler: %w", err)
		}
		admin.Administer(handler)

		ep := endpoint.Goa(genweb.NewEndpoints).Adapt(handler)

		transport := transport.GoaHTTP(httpweb.New, httpweb.Mount)
		httpSrv.RegisterHandler("/", transport.Adapt(ctx, ep, web.StaticFS))
	}

	err = app.New(httpSrv, grpcSrv, admin, worker).Run(ctx)
	if err != nil {
		ctxlog.Error(ctx, "encountered error while running app", err)
		return fmt.Errorf("app run: %w", err)
	}

	return nil
}
