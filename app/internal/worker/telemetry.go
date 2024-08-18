package worker

import (
	"context"
	"fmt"

	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"goa.design/clue/log"

	"github.com/jace-ys/countup/internal/instrument"
	"github.com/jace-ys/countup/internal/slog"
)

type poolMetrics struct {
	jobsEnqueuedTotal         metric.Int64Counter
	jobsCompletedTotal        metric.Int64Counter
	jobsFailedTotal           metric.Int64Counter
	jobsDiscardedTotal        metric.Int64Counter
	jobsCancelledTotal        metric.Int64Counter
	jobsAvailableCount        metric.Int64UpDownCounter
	jobsRunDurationSeconds    metric.Float64Histogram
	jobsQueueWaitMilliseconds metric.Int64Histogram
}

func (p *Pool) initMetrics() error {
	meter := instrument.OTel.Meter()

	metrics := new(poolMetrics)
	var err error

	metrics.jobsEnqueuedTotal, err = meter.Int64Counter("worker.jobs.enqueued.total")
	if err != nil {
		return fmt.Errorf("worker.jobs.enqueued.total: %w", err)
	}

	metrics.jobsCompletedTotal, err = meter.Int64Counter("worker.jobs.completed.total")
	if err != nil {
		return fmt.Errorf("worker.jobs.completed.total: %w", err)
	}

	metrics.jobsFailedTotal, err = meter.Int64Counter("worker.jobs.failed.total")
	if err != nil {
		return fmt.Errorf("worker.jobs.failed.total: %w", err)
	}

	metrics.jobsDiscardedTotal, err = meter.Int64Counter("worker.jobs.discarded.total")
	if err != nil {
		return fmt.Errorf("worker.jobs.discarded.total: %w", err)
	}

	metrics.jobsCancelledTotal, err = meter.Int64Counter("worker.jobs.cancelled.total")
	if err != nil {
		return fmt.Errorf("worker.jobs.cancelled.total: %w", err)
	}

	metrics.jobsAvailableCount, err = meter.Int64UpDownCounter("worker.jobs.available.count")
	if err != nil {
		return fmt.Errorf("worker.jobs.available.count: %w", err)
	}

	metrics.jobsRunDurationSeconds, err = meter.Float64Histogram("worker.jobs.run.duration.seconds")
	if err != nil {
		return fmt.Errorf("worker.jobs.run.duration.seconds: %w", err)
	}

	metrics.jobsQueueWaitMilliseconds, err = meter.Int64Histogram("worker.jobs.queue.wait.milliseconds")
	if err != nil {
		return fmt.Errorf("worker.jobs.queue.wait.milliseconds: %w", err)
	}

	p.metrics = metrics
	return nil
}

func (p *Pool) runMetricsExporter(ctx context.Context) {
	subscriptions := []river.EventKind{
		river.EventKindJobCompleted,
		river.EventKindJobCancelled,
		river.EventKindJobFailed,
	}

	events, close := p.pool.Subscribe(subscriptions...)
	defer close()

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-events:
			if !ok {
				return
			}

			attrs := attribute.NewSet([]attribute.KeyValue{
				attribute.String("job.worker", "river"),
				attribute.String("job.kind", event.Job.Kind),
				attribute.String("job.queue", event.Job.Queue),
				attribute.Int("job.priority", event.Job.Priority),
			}...)

			p.metrics.jobsQueueWaitMilliseconds.Record(ctx, event.JobStats.QueueWaitDuration.Milliseconds(), metric.WithAttributeSet(attrs))

			switch event.Kind {
			case river.EventKindJobCompleted:
				p.metrics.jobsCompletedTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))
				p.metrics.jobsAvailableCount.Add(ctx, -1, metric.WithAttributeSet(attrs))
				p.metrics.jobsRunDurationSeconds.Record(ctx, event.JobStats.RunDuration.Seconds(), metric.WithAttributeSet(attrs))

			case river.EventKindJobCancelled:
				p.metrics.jobsCancelledTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))
				p.metrics.jobsAvailableCount.Add(ctx, -1, metric.WithAttributeSet(attrs))

			case river.EventKindJobFailed:
				p.metrics.jobsFailedTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))

				switch event.Job.State {
				case rivertype.JobStateDiscarded:
					p.metrics.jobsDiscardedTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))
					p.metrics.jobsAvailableCount.Add(ctx, -1, metric.WithAttributeSet(attrs))
				}
			}
		}
	}
}

func (p *Pool) emitEnqueuedTelemetry(ctx context.Context, enqueued *rivertype.JobInsertResult) {
	kvs := []log.Fielder{
		slog.KV("job.worker", "river"),
		slog.KV("job.kind", enqueued.Job.Kind),
		slog.KV("job.id", enqueued.Job.ID),
		slog.KV("job.queue", enqueued.Job.Queue),
		slog.KV("job.priority", enqueued.Job.Priority),
		slog.KV("job.scheduled_at", enqueued.Job.ScheduledAt),
		slog.KV("job.max_attempts", enqueued.Job.MaxAttempts),
	}

	attrs := []attribute.KeyValue{
		attribute.String("job.worker", "river"),
		attribute.String("job.kind", enqueued.Job.Kind),
		attribute.String("job.queue", enqueued.Job.Queue),
		attribute.Int("job.priority", enqueued.Job.Priority),
	}

	slog.Print(ctx, "job enqueued", kvs...)

	span := trace.SpanFromContext(ctx)
	span.AddEvent("job.enqueued",
		trace.WithTimestamp(enqueued.Job.CreatedAt),
		trace.WithAttributes(attrs...),
		trace.WithAttributes(
			attribute.Int64("job.id", enqueued.Job.ID),
			attribute.String("job.scheduled_at", enqueued.Job.ScheduledAt.String()),
			attribute.Int("job.max_attempts", enqueued.Job.MaxAttempts),
		),
	)

	attrset := attribute.NewSet(attrs...)
	p.metrics.jobsEnqueuedTotal.Add(ctx, 1, metric.WithAttributeSet(attrset))
	p.metrics.jobsAvailableCount.Add(ctx, 1, metric.WithAttributeSet(attrset))
}
