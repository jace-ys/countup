package worker

import (
	"context"
	"fmt"

	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/jace-ys/countup/internal/instrument"
)

type metrics struct {
	jobsEnqueuedTotal         metric.Int64Counter
	jobsCompletedTotal        metric.Int64Counter
	jobsFailedTotal           metric.Int64Counter
	jobsDiscardedTotal        metric.Int64Counter
	jobsCancelledTotal        metric.Int64Counter
	jobsAvailableCount        metric.Int64UpDownCounter
	jobsRunDurationSeconds    metric.Float64Histogram
	jobsQueueWaitMilliseconds metric.Int64Histogram
}

func newMetrics() (*metrics, error) {
	meter := instrument.OTel.Meter()

	metrics := &metrics{}
	var err error

	metrics.jobsEnqueuedTotal, err = meter.Int64Counter("worker.jobs.enqueued.total")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.enqueued.total: %w", err)
	}

	metrics.jobsCompletedTotal, err = meter.Int64Counter("worker.jobs.completed.total")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.completed.total: %w", err)
	}

	metrics.jobsFailedTotal, err = meter.Int64Counter("worker.jobs.failed.total")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.failed.total: %w", err)
	}

	metrics.jobsDiscardedTotal, err = meter.Int64Counter("worker.jobs.discarded.total")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.discarded.total: %w", err)
	}

	metrics.jobsCancelledTotal, err = meter.Int64Counter("worker.jobs.cancelled.total")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.cancelled.total: %w", err)
	}

	metrics.jobsAvailableCount, err = meter.Int64UpDownCounter("worker.jobs.available.count")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.available.count: %w", err)
	}

	metrics.jobsRunDurationSeconds, err = meter.Float64Histogram("worker.jobs.run_duration.seconds")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.run.duration.seconds: %w", err)
	}

	metrics.jobsQueueWaitMilliseconds, err = meter.Int64Histogram("worker.jobs.queue_wait.milliseconds")
	if err != nil {
		return nil, fmt.Errorf("worker.jobs.queue.wait.milliseconds: %w", err)
	}

	return metrics, nil
}

func (p *Pool) runMetricsExporter(ctx context.Context) {
	subscriptions := []river.EventKind{
		river.EventKindJobCompleted,
		river.EventKindJobCancelled,
		river.EventKindJobFailed,
	}

	events, cancel := p.pool.Subscribe(subscriptions...)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-events:
			if !ok {
				return
			}

			attrs := attribute.NewSet(
				attribute.String("job.worker", "river"),
				attribute.String("job.kind", event.Job.Kind),
				attribute.String("job.queue", event.Job.Queue),
				attribute.Int("job.priority", event.Job.Priority),
			)

			p.metrics.jobsQueueWaitMilliseconds.Record(ctx, event.JobStats.QueueWaitDuration.Milliseconds(), metric.WithAttributeSet(attrs))

			switch event.Kind { //nolint:exhaustive
			case river.EventKindJobCompleted:
				p.metrics.jobsCompletedTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))
				p.metrics.jobsAvailableCount.Add(ctx, -1, metric.WithAttributeSet(attrs))
				p.metrics.jobsRunDurationSeconds.Record(ctx, event.JobStats.RunDuration.Seconds(), metric.WithAttributeSet(attrs))

			case river.EventKindJobCancelled:
				p.metrics.jobsCancelledTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))
				p.metrics.jobsAvailableCount.Add(ctx, -1, metric.WithAttributeSet(attrs))

			case river.EventKindJobFailed:
				p.metrics.jobsFailedTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))

				switch event.Job.State { //nolint:exhaustive
				case rivertype.JobStateDiscarded:
					p.metrics.jobsDiscardedTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))
					p.metrics.jobsAvailableCount.Add(ctx, -1, metric.WithAttributeSet(attrs))
				}
			}
		}
	}
}
