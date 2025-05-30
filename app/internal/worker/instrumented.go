package worker

import (
	"context"
	"errors"
	"fmt"

	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"goa.design/clue/log"

	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/instrument"
)

type instrumentedWorkerMiddleware struct {
	river.WorkerMiddlewareDefaults
}

var _ rivertype.WorkerMiddleware = (*instrumentedWorkerMiddleware)(nil)

func (w *instrumentedWorkerMiddleware) Work(ctx context.Context, job *rivertype.JobRow, doInner func(context.Context) error) error {
	kvs := []log.Fielder{
		ctxlog.KV("job.worker", "river"),
		ctxlog.KV("job.kind", job.Kind),
		ctxlog.KV("job.id", job.ID),
		ctxlog.KV("job.attempt", job.Attempt),
	}

	attrs := []attribute.KeyValue{
		attribute.String("job.worker", "river"),
		attribute.String("job.kind", job.Kind),
		attribute.Int64("job.id", job.ID),
		attribute.Int("job.attempt", job.Attempt),
		attribute.String("job.queue", job.Queue),
		attribute.Int("job.priority", job.Priority),
		attribute.String("job.scheduled_at", job.ScheduledAt.String()),
		attribute.String("job.attempted_at", job.AttemptedAt.String()),
	}

	md, err := parseMetadata(job.Metadata)
	if err != nil {
		ctxlog.Error(ctx, "failed to extract metadata from job", err)
	}

	for k, v := range md {
		kvs = append(kvs, ctxlog.KV(k, v))
		attrs = append(attrs, attribute.String(k, v))
	}

	parentSpanCtx := instrument.OTel.Propagators().Extract(context.Background(), md)

	source := fmt.Sprintf("river.worker/%s", job.Kind)
	ctx, span := instrument.OTel.Tracer().Start(ctx, source, trace.WithAttributes(attrs...))
	span.AddLink(trace.LinkFromContext(parentSpanCtx))
	defer span.End()

	ctx = ctxlog.With(ctx, kvs...)
	ctxlog.Print(ctx, "job started",
		ctxlog.KV("job.queue", job.Queue),
		ctxlog.KV("job.priority", job.Priority),
		ctxlog.KV("job.scheduled_at", job.ScheduledAt),
		ctxlog.KV("job.attempted_at", job.AttemptedAt),
	)

	var jobErr error
	func() {
		defer func() {
			if rvr := recover(); rvr != nil {
				instrument.EmitRecoveredPanicTelemetry(ctx, rvr, source)
				jobErr = fmt.Errorf("%v", rvr)
			}
		}()
		jobErr = doInner(ctx)
	}()

	if jobErr != nil {
		var jobErrReason string
		switch {
		case errors.Is(err, &river.JobSnoozeError{}):
			ctxlog.Print(ctx, "job snoozed")
			span.SetAttributes(attribute.Bool("job.snoozed", true))
			return jobErr

		case errors.Is(err, &river.JobCancelError{}):
			jobErrReason = "job cancelled"
			span.SetAttributes(attribute.Bool("job.cancelled", true))

		case job.Attempt == job.MaxAttempts:
			jobErrReason = "job failed, discarded due to max attempts exceeded"
			span.SetAttributes(attribute.Bool("job.discarded", true))

		default:
			jobErrReason = "job failed"
			span.SetAttributes(attribute.Bool("job.failed", true))
		}

		ctxlog.Error(ctx, jobErrReason, jobErr)
		span.SetStatus(codes.Error, jobErrReason)
		span.SetAttributes(attribute.String("error", jobErr.Error()))

		return jobErr
	}

	ctxlog.Print(ctx, "job completed")
	return nil
}

type instrumentedJobInsertMiddleware struct {
	river.MiddlewareDefaults
	metrics *metrics
}

var _ rivertype.Middleware = (*instrumentedJobInsertMiddleware)(nil)

func (m *instrumentedJobInsertMiddleware) InsertMany(ctx context.Context, manyParams []*rivertype.JobInsertParams, doInner func(ctx context.Context) ([]*rivertype.JobInsertResult, error)) ([]*rivertype.JobInsertResult, error) {
	results, err := doInner(ctx)
	for _, enqueued := range results {
		m.emitEnqueuedTelemetry(ctx, enqueued)
	}
	return results, err
}

func (m *instrumentedJobInsertMiddleware) emitEnqueuedTelemetry(ctx context.Context, enqueued *rivertype.JobInsertResult) {
	kvs := []log.Fielder{
		ctxlog.KV("job.worker", "river"),
		ctxlog.KV("job.kind", enqueued.Job.Kind),
		ctxlog.KV("job.id", enqueued.Job.ID),
		ctxlog.KV("job.queue", enqueued.Job.Queue),
		ctxlog.KV("job.priority", enqueued.Job.Priority),
		ctxlog.KV("job.scheduled_at", enqueued.Job.ScheduledAt),
		ctxlog.KV("job.max_attempts", enqueued.Job.MaxAttempts),
	}

	attrs := []attribute.KeyValue{
		attribute.String("job.worker", "river"),
		attribute.String("job.kind", enqueued.Job.Kind),
		attribute.String("job.queue", enqueued.Job.Queue),
		attribute.Int("job.priority", enqueued.Job.Priority),
	}

	ctxlog.Print(ctx, "job enqueued", kvs...)

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

	attrSet := attribute.NewSet(attrs...)
	m.metrics.jobsEnqueuedTotal.Add(ctx, 1, metric.WithAttributeSet(attrSet))
	m.metrics.jobsAvailableCount.Add(ctx, 1, metric.WithAttributeSet(attrSet))
}
