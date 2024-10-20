package worker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/alexliesenfeld/health"
	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivertype"

	"github.com/jace-ys/countup/internal/app"
	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/healthz"
	"github.com/jace-ys/countup/internal/postgres"
)

func Register[T river.JobArgs](pool *Pool, worker river.Worker[T]) {
	river.AddWorker(pool.workers, worker)
}

type Pool struct {
	name    string
	pool    *river.Client[pgx.Tx]
	workers *river.Workers
	metrics *metrics
}

func NewPool(ctx context.Context, name string, db *postgres.Pool, w io.Writer, concurrency int) (*Pool, error) {
	metrics, err := newMetrics()
	if err != nil {
		return nil, fmt.Errorf("init metrics: %w", err)
	}

	workers := river.NewWorkers()

	client, err := river.NewClient(riverpgxv5.New(db.Pool), &river.Config{
		Workers: workers,
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: concurrency},
		},
		JobInsertMiddleware: []rivertype.JobInsertMiddleware{
			&instrumentedJobInsertMiddleware{metrics: metrics},
		},
		WorkerMiddleware: []rivertype.WorkerMiddleware{
			&instrumentedWorkerMiddleware{},
		},
		Logger: slog.New(ctxlog.AsNopHandler()),
	})
	if err != nil {
		return nil, fmt.Errorf("init river client: %w", err)
	}

	return &Pool{
		name:    name,
		pool:    client,
		workers: workers,
		metrics: metrics,
	}, nil
}

func (p *Pool) Enqueue(ctx context.Context, job river.JobArgs, opts ...EnqueueOption) error {
	opts = append(opts, withContextMetadata(ctx))

	insertOpts := &river.InsertOpts{}
	for _, opt := range opts {
		opt(insertOpts)
	}

	_, err := p.pool.Insert(ctx, job, insertOpts)
	return err //nolint:wrapcheck
}

func (p *Pool) EnqueueTx(ctx context.Context, tx pgx.Tx, job river.JobArgs, opts ...EnqueueOption) error {
	opts = append(opts, withContextMetadata(ctx))

	insertOpts := &river.InsertOpts{}
	for _, opt := range opts {
		opt(insertOpts)
	}

	_, err := p.pool.InsertTx(ctx, tx, job, insertOpts)
	return err //nolint:wrapcheck
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
	return p.pool.StopAndCancel(ctx) //nolint:wrapcheck
}

var _ healthz.Target = (*Pool)(nil)

func (p *Pool) HealthChecks() []health.Check {
	return []health.Check{
		{
			Name: fmt.Sprintf("%s:%s", p.Kind(), p.Name()),
			Check: func(ctx context.Context) error {
				select {
				case <-p.pool.Stopped():
					return errors.New("river client reported as not running")
				default:
					return nil
				}
			},
		},
	}
}
