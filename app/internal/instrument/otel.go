package instrument

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
	"goa.design/clue/clue"

	"github.com/jace-ys/countup/internal/versioninfo"
)

var OTel *OTelProvider

type OTelProvider struct {
	cfg           *clue.Config
	shutdownFuncs []func(context.Context) error
}

func MustInitOTelProvider(ctx context.Context, name, version, otlpMetricsEndpoint, otlpTracesEndpoint string) {
	if err := InitOTelProvider(ctx, name, version, otlpMetricsEndpoint, otlpTracesEndpoint); err != nil {
		panic(err)
	}
}

func InitOTelProvider(ctx context.Context, name, version, otlpMetricsEndpoint, otlpTracesEndpoint string) error {
	var shutdownFuncs []func(context.Context) error

	metrics, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(otlpMetricsEndpoint),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return fmt.Errorf("init metrics exporter: %w", err)
	}
	shutdownFuncs = append(shutdownFuncs, metrics.Shutdown)

	traces, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(otlpTracesEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return fmt.Errorf("init trace exporter: %w", err)
	}
	shutdownFuncs = append(shutdownFuncs, traces.Shutdown)

	opts := []clue.Option{
		clue.WithResource(resource.Environment()),
		clue.WithReaderInterval(30 * time.Second),
	}

	cfg, err := clue.NewConfig(ctx, name, version, metrics, traces, opts...)
	if err != nil {
		return fmt.Errorf("init otel provider: %w", err)
	}
	clue.ConfigureOpenTelemetry(ctx, cfg)

	OTel = &OTelProvider{
		cfg:           cfg,
		shutdownFuncs: shutdownFuncs,
	}

	if err := OTel.initMetrics(ctx); err != nil {
		return fmt.Errorf("init metrics: %w", err)
	}

	return nil
}

func (i *OTelProvider) initMetrics(ctx context.Context) error {
	if err := runtime.Start(); err != nil {
		return fmt.Errorf("runtime: %w", err)
	}

	up, err := OTel.Meter().Int64Gauge("up")
	if err != nil {
		return fmt.Errorf("up: %w", err)
	}
	up.Record(ctx, 1)

	return nil
}

func (i *OTelProvider) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	select {
	case <-time.After(5 * time.Second):
	case <-ctx.Done():
		return ctx.Err() //nolint:wrapcheck
	}

	var err error
	for _, fn := range i.shutdownFuncs {
		err = errors.Join(err, fn(ctx))
	}

	i.shutdownFuncs = nil
	return err
}

const scope = "github.com/jace-ys/countup/internal/instrument"

func (i *OTelProvider) Meter() metric.Meter {
	return i.cfg.MeterProvider.Meter(scope, metric.WithInstrumentationVersion(versioninfo.Version))
}

func (i *OTelProvider) Tracer() trace.Tracer {
	return i.cfg.TracerProvider.Tracer(scope, trace.WithInstrumentationVersion(versioninfo.Version))
}

func (i *OTelProvider) Propagators() propagation.TextMapPropagator {
	return i.cfg.Propagators
}
