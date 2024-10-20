package score

import (
	"context"
	"fmt"

	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/postgres"
	scorestore "github.com/jace-ys/countup/internal/service/score/store"
)

type Service struct {
	db    *postgres.Pool
	store scorestore.Querier
}

func New(db *postgres.Pool, store scorestore.Querier) *Service {
	return &Service{
		db:    db,
		store: store,
	}
}

func (s *Service) ListScores(ctx context.Context) ([]Score, error) {
	ctxlog.Info(ctx, "listing scores")

	scores, err := s.store.ListScores(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrListScores, err)
	}

	res := make([]Score, len(scores))
	for i, v := range scores {
		res[i] = Score{
			UserEmail: v.UserEmail,
			Score:     v.Score,
		}
	}

	return res, nil
}

func (s *Service) InsertScore(ctx context.Context, userEmail string, score int32) error {
	ctxlog.Info(ctx, "inserting score")

	if err := s.store.InsertScore(ctx, s.db, scorestore.InsertScoreParams{
		UserEmail: userEmail,
		Score:     score,
	}); err != nil {
		return fmt.Errorf("%w: %w", ErrInsertScore, err)
	}

	return nil
}
