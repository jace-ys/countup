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
	goose.AddMigrationNoTxContext(upRiverMigrate006, downRiverMigrate006)
}

func upRiverMigrate006(ctx context.Context, db *sql.DB) error {
	migrator, err := rivermigrate.New(riverdatabasesql.New(db), &rivermigrate.Config{})
	if err != nil {
		return fmt.Errorf("init river migrator: %w", err)
	}

	_, err = migrator.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 6,
	})
	if err != nil {
		return fmt.Errorf("apply river migration: %w", err)
	}

	return nil
}

func downRiverMigrate006(ctx context.Context, db *sql.DB) error {
	migrator, err := rivermigrate.New(riverdatabasesql.New(db), &rivermigrate.Config{})
	if err != nil {
		return fmt.Errorf("init river migrator: %w", err)
	}

	_, err = migrator.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		TargetVersion: 5,
	})
	if err != nil {
		return fmt.Errorf("apply river migration: %w", err)
	}

	return nil
}
