package counter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	counterstore "github.com/jace-ys/countup/internal/service/counter/store"
	"github.com/jace-ys/countup/internal/slog"
	"github.com/jace-ys/countup/internal/worker"
)

type Service struct {
	db      *pgxpool.Pool
	workers *worker.Pool
	store   counterstore.Querier

	finalizeWindow time.Duration
}

func New(db *pgxpool.Pool, workers *worker.Pool, store counterstore.Querier, finalizeWindow time.Duration) *Service {
	worker.Register(workers, &FinalizeIncrementWorker{
		db:    db,
		store: store,
	})

	return &Service{
		db:             db,
		workers:        workers,
		store:          store,
		finalizeWindow: finalizeWindow,
	}
}

func (s *Service) GetInfo(ctx context.Context) (*Info, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: tx begin: %w", ErrDBConn, err)
	}
	defer tx.Rollback(ctx)

	slog.Info(ctx, "getting counter")

	counter, err := s.store.GetCounter(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetCounter, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("%w: tx commit: %w", ErrDBConn, err)
	}

	return &Info{
		Count:           counter.Count,
		LastIncrementBy: counter.LastIncrementBy.String,
		LastIncrementAt: counter.LastIncrementAt.Time,
		NextFinalizeAt:  counter.NextFinalizeAt.Time,
	}, nil
}

func (s *Service) RequestIncrement(ctx context.Context, user string) error {
	ctx = slog.With(ctx, slog.KV("request.user", user))

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%w: tx begin: %w", ErrDBConn, err)
	}
	defer tx.Rollback(ctx)

	slog.Info(ctx, "inserting increment request")

	if err := s.store.InsertIncrementRequest(ctx, tx, counterstore.InsertIncrementRequestParams{
		RequestedBy: user,
		RequestedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return &MultipleIncrementRequestError{user, s.finalizeWindow}
			}
		}
		return fmt.Errorf("%w: %w", ErrInsertIncrementRequest, err)
	}

	slog.Info(ctx, "listing increment request")

	requests, err := s.store.ListIncrementRequests(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrListIncrementRequests, err)
	}

	if len(requests) > 1 {
		slog.Info(ctx, "existing increment requests in finalize window, skip enqueuing finalize job",
			slog.KV("requests.count", len(requests)),
			slog.KV("finalize.window", s.finalizeWindow),
		)
	} else {
		slog.Info(ctx, "first increment request in finalize window, enqueuing finalize job",
			slog.KV("finalize.window", s.finalizeWindow),
		)

		finalizeAt := time.Now().Add(s.finalizeWindow)

		if err := s.workers.EnqueueTx(ctx, tx, FinalizeIncrementJobArgs{
			FinalizeWindow: s.finalizeWindow,
		},
			worker.WithSchedule(finalizeAt),
		); err != nil {
			return fmt.Errorf("%w: %w", ErrEnqueueFinalizeIncrement, err)
		}

		slog.Info(ctx, "updating counter finalize time", slog.KV("finalize.at", finalizeAt))

		if err := s.store.UpdateCounterFinalizeTime(ctx, tx, pgtype.Timestamptz{
			Time:  finalizeAt,
			Valid: true,
		}); err != nil {
			return fmt.Errorf("%w: %w", ErrUpdateCounterFinalizeTime, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("%w: tx commit: %w", ErrDBConn, err)
	}

	return nil
}
