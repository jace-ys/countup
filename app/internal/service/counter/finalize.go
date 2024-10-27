package counter

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"

	counterstore "github.com/jace-ys/countup/internal/service/counter/store"
	"github.com/jace-ys/countup/internal/slog"
)

type FinalizeIncrementJobArgs struct {
	FinalizeWindow time.Duration
}

func (FinalizeIncrementJobArgs) Kind() string {
	return "counter.FinalizeIncrement"
}

type FinalizeIncrementWorker struct {
	river.WorkerDefaults[FinalizeIncrementJobArgs]

	db    *pgxpool.Pool
	store counterstore.Querier
}

func NewIncrementWorker(db *pgxpool.Pool, store counterstore.Querier) *FinalizeIncrementWorker {
	return &FinalizeIncrementWorker{
		db:    db,
		store: store,
	}
}

func (w *FinalizeIncrementWorker) Work(ctx context.Context, job *river.Job[FinalizeIncrementJobArgs]) error {
	ctx = slog.With(ctx, slog.KV("finalize.window", job.Args.FinalizeWindow))

	tx, err := w.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%w: tx begin: %w", ErrDBConn, err)
	}
	defer tx.Rollback(ctx)

	slog.Info(ctx, "listing increment requests")

	requests, err := w.store.ListIncrementRequests(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrListIncrementRequests, err)
	}

	switch len(requests) {
	case 0:
		slog.Info(ctx, "no increment requests in finalize window, returning",
			slog.KV("finalize.window", job.Args.FinalizeWindow),
		)
		return nil

	case 1:
		slog.Info(ctx, "only one increment request in finalize window, incrementing counter",
			slog.KV("finalize.user", requests[0].RequestedBy),
		)

		if err := w.store.IncrementCounter(ctx, tx, counterstore.IncrementCounterParams{
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
		slog.Info(ctx, "multiple increment requests in finalize window, resetting counter",
			slog.KV("requests.count", len(requests)),
			slog.KV("finalize.user", requests[0].RequestedBy),
		)

		if err := w.store.ResetCounter(ctx, tx); err != nil {
			return fmt.Errorf("%w: %w", ErrResetCounter, err)
		}
	}

	slog.Info(ctx, "deleting increment requests")

	if err := w.store.DeleteIncrementRequests(ctx, tx); err != nil {
		return fmt.Errorf("%w: %w", ErrTruncateIncrementRequests, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("%w: tx commit: %w", ErrDBConn, err)
	}

	return nil
}
