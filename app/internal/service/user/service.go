package user

import (
	"context"
	"fmt"

	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/idgen"
	"github.com/jace-ys/countup/internal/postgres"
	userstore "github.com/jace-ys/countup/internal/service/user/store"
)

type Service struct {
	db    *postgres.Pool
	store userstore.Querier
}

func New(db *postgres.Pool, store userstore.Querier) *Service {
	return &Service{
		db:    db,
		store: store,
	}
}

func (s *Service) GetUser(ctx context.Context, userID idgen.ID[idgen.User]) (*User, error) {
	ctxlog.Info(ctx, "getting user", ctxlog.KV("user.id", userID))

	user, err := s.store.GetUser(ctx, s.db, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUser, err)
	}

	return &User{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (s *Service) CreateUserIfNotExists(ctx context.Context, email string) (*User, error) {
	ctxlog.Info(ctx, "creating user if not exists", ctxlog.KV("user.email", email))

	if err := s.store.InsertUserIfNotExists(ctx, s.db, userstore.InsertUserIfNotExistsParams{
		ID:    idgen.New[idgen.User](),
		Email: email,
	}); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInsertUserIfNotExists, err)
	}

	ctxlog.Info(ctx, "getting user by email", ctxlog.KV("user.email", email))

	user, err := s.store.GetUserByEmail(ctx, s.db, email)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetUser, err)
	}

	return &User{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
