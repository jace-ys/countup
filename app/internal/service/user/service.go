package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jace-ys/countup/internal/idgen"
	userstore "github.com/jace-ys/countup/internal/service/user/store"
	"github.com/jace-ys/countup/internal/slog"
)

type Service struct {
	db    *pgxpool.Pool
	store userstore.Querier
}

func New(db *pgxpool.Pool, store userstore.Querier) *Service {
	return &Service{
		db:    db,
		store: store,
	}
}

func (s *Service) GetUser(ctx context.Context, userID string) (*User, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	slog.Info(ctx, "getting user", slog.KV("user.id", userID))

	user, err := s.store.GetUser(ctx, tx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by ID: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &User{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (s *Service) CreateUserIfNotExists(ctx context.Context, email string) (*User, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	slog.Info(ctx, "creating user if not exists", slog.KV("user.email", email))

	user, err := s.store.InsertUserIfNotExists(ctx, tx, userstore.InsertUserIfNotExistsParams{
		ID:    idgen.NewID("usr"),
		Email: email,
	})
	if err != nil {
		return nil, fmt.Errorf("insert user if not exists: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &User{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
