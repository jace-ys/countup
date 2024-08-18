package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/attribute"
)

func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	attrs := []attribute.KeyValue{
		attribute.String("db.database", cfg.ConnConfig.Database),
	}

	cfg.ConnConfig.Tracer = otelpgx.NewTracer(
		otelpgx.WithAttributes(attrs...),
		otelpgx.WithTrimSQLInSpanName(),
		otelpgx.WithSpanNameFunc(func(stmt string) string {
			idx := strings.IndexRune(stmt, '\n')
			if idx >= 0 {
				return stmt[:idx]
			}
			return stmt
		}),
	)

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	return pool, nil
}
