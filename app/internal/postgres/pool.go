package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/attribute"
)

type Pool struct {
	*pgxpool.Pool
}

func NewPool(ctx context.Context, connString string) (*Pool, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	attrs := []attribute.KeyValue{
		attribute.String("db.database", cfg.ConnConfig.Database),
		attribute.String("db.user", cfg.ConnConfig.User),
	}

	cfg.ConnConfig.Tracer = otelpgx.NewTracer(
		otelpgx.WithTracerAttributes(attrs...),
		otelpgx.WithMeterAttributes(attrs...),
		otelpgx.WithTrimSQLInSpanName(),
		otelpgx.WithSpanNameFunc(func(stmt string) string {
			idx := strings.IndexRune(stmt, '\n')
			if idx >= 0 {
				return stmt[:idx]
			}
			return stmt
		}),
	)

	cfg.ConnConfig.ConnectTimeout = 10 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	if err := otelpgx.RecordStats(pool, otelpgx.WithStatsAttributes(attrs...)); err != nil {
		return nil, fmt.Errorf("record stats: %w", err)
	}

	return &Pool{pool}, nil
}

func (p *Pool) WithinTx(ctx context.Context, fn func(tx pgx.Tx) error) error {
	return pgx.BeginFunc(ctx, p, fn) //nolint:wrapcheck
}
