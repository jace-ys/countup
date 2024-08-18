package worker

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/riverqueue/river"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"goa.design/clue/log"

	"github.com/jace-ys/countup/internal/instrument"
	"github.com/jace-ys/countup/internal/slog"
)

type instrumentedWorker[T river.JobArgs] struct {
	worker river.Worker[T]
}

func (w *instrumentedWorker[T]) NextRetry(job *river.Job[T]) time.Time {
	return w.worker.NextRetry(job)
}

func (w *instrumentedWorker[T]) Timeout(job *river.Job[T]) time.Duration {
	return w.worker.Timeout(job)
}

func (w *instrumentedWorker[T]) Work(ctx context.Context, job *river.Job[T]) (err error) {
	kvs := []log.Fielder{
		slog.KV("job.worker", "river"),
		slog.KV("job.kind", job.Kind),
		slog.KV("job.id", job.ID),
		slog.KV("job.attempt", job.Attempt),
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
		slog.Error(ctx, "failed to extract metadata from job", err)
	}

	for k, v := range md {
		kvs = append(kvs, slog.KV(k, v))
		attrs = append(attrs, attribute.String(k, v))
	}

	source := fmt.Sprintf("river.worker/%s", job.Kind)
	ctx, span := instrument.OTel.Tracer().Start(ctx, source)
	span.SetAttributes(attrs...)
	defer span.End()

	ctx = slog.With(ctx, kvs...)
	slog.Print(ctx, "job started",
		slog.KV("job.queue", job.Queue),
		slog.KV("job.priority", job.Priority),
		slog.KV("job.scheduled_at", job.ScheduledAt),
		slog.KV("job.attempted_at", job.AttemptedAt),
	)

	func() {
		defer func() {
			if rvr := recover(); rvr != nil {
				err = instrument.EmitRecoveredPanicTelemetry(ctx, rvr, source)
			}
		}()
		err = w.worker.Work(ctx, job)
	}()

	if err != nil {
		var errReason string
		switch {
		case strings.HasPrefix(err.Error(), "jobCancelError:"):
			errReason = "job cancelled"
			span.SetAttributes(attribute.Bool("job.cancelled", true))

		case strings.HasPrefix(err.Error(), "jobSnoozeError:"):
			slog.Print(ctx, "job snoozed")
			span.SetAttributes(attribute.Bool("job.snoozed", true))
			return err

		case job.Attempt == job.MaxAttempts:
			errReason = "job failed, discarded due to max attempts exceeded"
			span.SetAttributes(attribute.Bool("job.discarded", true))

		default:
			errReason = "job failed"
			span.SetAttributes(attribute.Bool("job.failed", true))
		}

		slog.Error(ctx, errReason, err)
		span.SetStatus(codes.Error, errReason)
		span.SetAttributes(attribute.String("error", err.Error()))

		return err
	}

	slog.Print(ctx, "job completed")
	return nil
}
