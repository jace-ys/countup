package counter

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/riverqueue/river"

	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/postgres"
	counterstore "github.com/jace-ys/countup/internal/service/counter/store"
)

type FinalizeIncrementJobArgs struct {
	FinalizeWindow time.Duration
}

func (FinalizeIncrementJobArgs) Kind() string {
	return "counter.FinalizeIncrement"
}

type FinalizeIncrementWorker struct {
	river.WorkerDefaults[FinalizeIncrementJobArgs]

	db    *postgres.Pool
	store counterstore.Querier
}

func NewIncrementWorker(db *postgres.Pool, store counterstore.Querier) *FinalizeIncrementWorker {
	return &FinalizeIncrementWorker{
		db:    db,
		store: store,
	}
}

func (w *FinalizeIncrementWorker) Work(ctx context.Context, job *river.Job[FinalizeIncrementJobArgs]) error {
	ctx = ctxlog.With(ctx, ctxlog.KV("finalize.window", job.Args.FinalizeWindow))

	ctxlog.Info(ctx, "listing increment requests")

	requests, err := w.store.ListIncrementRequests(ctx, w.db)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrListIncrementRequests, err)
	}

	switch len(requests) {
	case 0:
		ctxlog.Info(ctx, "no increment requests in finalize window, returning")
		return nil

	case 1:
		ctxlog.Info(ctx, "only one increment request in finalize window, incrementing counter",
			ctxlog.KV("finalize.user", requests[0].RequestedBy),
		)

		if err := w.store.IncrementCounter(ctx, w.db, counterstore.IncrementCounterParams{
			LastIncrementBy: pgtype.Text{
				String: requests[0].RequestedBy,
				Valid:  true,
			},
			LastIncrementAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
		}); err != nil {
			return fmt.Errorf("%w: %w", ErrIncrementCounter, err)
		}

	default:
		ctxlog.Info(ctx, "multiple increment requests in finalize window, resetting counter",
			ctxlog.KV("requests.count", len(requests)),
			ctxlog.KV("finalize.user", requests[0].RequestedBy),
		)

		if err := w.store.ResetCounter(ctx, w.db); err != nil {
			return fmt.Errorf("%w: %w", ErrResetCounter, err)
		}
	}

	ctxlog.Info(ctx, "deleting increment requests")

	if err := w.store.DeleteIncrementRequests(ctx, w.db); err != nil {
		return fmt.Errorf("%w: %w", ErrTruncateIncrementRequests, err)
	}

	return nil
}
