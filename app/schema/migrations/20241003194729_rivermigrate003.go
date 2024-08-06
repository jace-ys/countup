package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/riverqueue/river/rivermigrate"
)

func init() {
	goose.AddMigrationContext(upRiverMigrate003, downRiverMigrate003)
}

func upRiverMigrate003(ctx context.Context, tx *sql.Tx) error {
	_, err := rivermigrator.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 3,
	})
	return err
}

func downRiverMigrate003(ctx context.Context, tx *sql.Tx) error {
	_, err := rivermigrator.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		TargetVersion: 2,
	})
	return err
}
