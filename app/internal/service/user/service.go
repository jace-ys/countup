package user

import (
	"github.com/jackc/pgx/v5/pgxpool"

	counterstore "github.com/jace-ys/countup/internal/service/counter/store"
)

type Service struct {
	db    *pgxpool.Pool
	store counterstore.Querier
}

func New(db *pgxpool.Pool, store counterstore.Querier) *Service {
	return &Service{
		db:    db,
		store: store,
	}
}
