package counter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/postgres"
	counterstore "github.com/jace-ys/countup/internal/service/counter/store"
	"github.com/jace-ys/countup/internal/worker"
)

type Service struct {
	db      *postgres.Pool
	workers *worker.Pool
	store   counterstore.Querier

	finalizeWindow time.Duration
}

func New(db *postgres.Pool, workers *worker.Pool, store counterstore.Querier, finalizeWindow time.Duration) *Service {
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
	ctxlog.Info(ctx, "getting counter")

	counter, err := s.store.GetCounter(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetCounter, err)
	}

	return &Info{
		Count:           counter.Count,
		LastIncrementBy: counter.LastIncrementBy.String,
		LastIncrementAt: counter.LastIncrementAt.Time,
		NextFinalizeAt:  counter.NextFinalizeAt.Time,
	}, nil
}

func (s *Service) RequestIncrement(ctx context.Context, user string) error {
	return s.db.WithinTx(ctx, func(tx pgx.Tx) error {
		ctxlog.Info(ctx, "inserting increment request")

		reqCount, err := s.store.InsertIncrementRequest(ctx, tx, counterstore.InsertIncrementRequestParams{
			RequestedBy: user,
			RequestedAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
		})
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				switch pgErr.Code {
				case pgerrcode.UniqueViolation:
					return &MultipleIncrementRequestError{user, s.finalizeWindow}
				}
			}
			return fmt.Errorf("%w: %w", ErrInsertIncrementRequest, err)
		}

		if reqCount > 1 {
			ctxlog.Info(ctx, "existing increment requests in finalize window, skip enqueuing finalize job",
				ctxlog.KV("requests.count", reqCount),
				ctxlog.KV("finalize.window", s.finalizeWindow),
			)
			return nil
		}

		ctxlog.Info(ctx, "first increment request in finalize window, enqueuing finalize job",
			ctxlog.KV("finalize.window", s.finalizeWindow),
		)

		finalizeAt := time.Now().Add(s.finalizeWindow)

		if err := s.workers.EnqueueTx(ctx, tx, FinalizeIncrementJobArgs{
			FinalizeWindow: s.finalizeWindow,
		},
			worker.WithSchedule(finalizeAt),
		); err != nil {
			return fmt.Errorf("%w: %w", ErrEnqueueFinalizeIncrement, err)
		}

		ctxlog.Info(ctx, "updating counter finalize time", ctxlog.KV("finalize.at", finalizeAt))

		if err := s.store.UpdateCounterFinalizeTime(ctx, tx, pgtype.Timestamptz{
			Time:  finalizeAt,
			Valid: true,
		}); err != nil {
			return fmt.Errorf("%w: %w", ErrUpdateCounterFinalizeTime, err)
		}

		return nil
	})
}
