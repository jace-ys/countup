package worker

import (
	"context"
	"errors"
	"fmt"
	"io"
	stdslog "log/slog"

	"github.com/alexliesenfeld/health"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"

	"github.com/jace-ys/countup/internal/app"
	"github.com/jace-ys/countup/internal/healthz"
	"github.com/jace-ys/countup/internal/slog"
)

func Register[T river.JobArgs](pool *Pool, worker river.Worker[T]) {
	river.AddWorker(pool.workers, &instrumentedWorker[T]{worker})
}

type Pool struct {
	name    string
	pool    *river.Client[pgx.Tx]
	workers *river.Workers

	metrics *poolMetrics
}

func NewPool(ctx context.Context, name string, db *pgxpool.Pool, w io.Writer, size int) (*Pool, error) {
	workers := river.NewWorkers()

	client, err := river.NewClient(riverpgxv5.New(db), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: size},
		},
		Workers: workers,
		Logger:  stdslog.New(slog.AsNopHandler(ctx)),
	})
	if err != nil {
		return nil, fmt.Errorf("init river client: %w", err)
	}

	pool := &Pool{
		name:    name,
		pool:    client,
		workers: workers,
	}

	if err := pool.initMetrics(); err != nil {
		return nil, fmt.Errorf("init metrics: %w", err)
	}

	return pool, nil
}

type EnqueueOpts func(*river.InsertOpts)

type EnqueuePostCommitHook func()

func (p *Pool) Enqueue(ctx context.Context, job river.JobArgs, opts ...EnqueueOpts) error {
	opts = append(opts, withContextMetadata(ctx))

	insertOpts := &river.InsertOpts{}
	for _, opt := range opts {
		opt(insertOpts)
	}

	enqueued, err := p.pool.Insert(ctx, job, insertOpts)
	if err != nil {
		return err
	}

	p.emitEnqueuedTelemetry(ctx, enqueued)
	return nil
}

func (p *Pool) EnqueueTx(ctx context.Context, tx pgx.Tx, job river.JobArgs, opts ...EnqueueOpts) (EnqueuePostCommitHook, error) {
	opts = append(opts, withContextMetadata(ctx))

	insertOpts := &river.InsertOpts{}
	for _, opt := range opts {
		opt(insertOpts)
	}

	enqueued, err := p.pool.InsertTx(ctx, tx, job, insertOpts)
	if err != nil {
		return nil, err
	}

	return func() { p.emitEnqueuedTelemetry(ctx, enqueued) }, nil
}

var _ app.Server = (*Pool)(nil)

func (p *Pool) Name() string {
	return p.name
}

func (p *Pool) Kind() string {
	return "worker"
}

func (p *Pool) Addr() string {
	return ""
}

func (p *Pool) Serve(ctx context.Context) error {
	go p.runMetricsExporter(ctx)
	if err := p.pool.Start(ctx); err != nil {
		return fmt.Errorf("starting worker pool: %w", err)
	}
	return nil
}

func (p *Pool) Shutdown(ctx context.Context) error {
	return p.pool.StopAndCancel(ctx)
}

var _ healthz.Target = (*Pool)(nil)

func (s *Pool) HealthChecks() []health.Check {
	return []health.Check{
		{
			Name: fmt.Sprintf("%s:%s", s.Kind(), s.Name()),
			Check: func(ctx context.Context) error {
				select {
				case <-s.pool.Stopped():
					return errors.New("river client reported as not running")
				default:
					return nil
				}
			},
		},
	}
}
