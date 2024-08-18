package migrations

import (
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivermigrate"
)

var rivermigrator *rivermigrate.Migrator[pgx.Tx]

func WithRiverMigrate(db *pgxpool.Pool) error {
	migrator, err := rivermigrate.New(riverpgxv5.New(db), &rivermigrate.Config{})
	if err != nil {
		return fmt.Errorf("init river migrator: %w", err)
	}

	rivermigrator = migrator
	return nil
}
