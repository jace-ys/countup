package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/riverqueue/river/riverdriver/riverdatabasesql"
	"github.com/riverqueue/river/rivermigrate"
)

func init() {
	goose.AddMigrationNoTxContext(upRiverMigrate003, downRiverMigrate003)
}

func upRiverMigrate003(ctx context.Context, db *sql.DB) error {
	migrator, err := rivermigrate.New(riverdatabasesql.New(db), &rivermigrate.Config{})
	if err != nil {
		return fmt.Errorf("init river migrator: %w", err)
	}

	_, err = migrator.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 3,
	})
	if err != nil {
		return fmt.Errorf("apply river migration: %w", err)
	}

	return nil
}

func downRiverMigrate003(ctx context.Context, db *sql.DB) error {
	migrator, err := rivermigrate.New(riverdatabasesql.New(db), &rivermigrate.Config{})
	if err != nil {
		return fmt.Errorf("init river migrator: %w", err)
	}

	_, err = migrator.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		TargetVersion: 2,
	})
	if err != nil {
		return fmt.Errorf("apply river migration: %w", err)
	}

	return nil
}
