package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/riverqueue/river/rivermigrate"
)

func init() {
	goose.AddMigrationContext(upRiverMigrate006, downRiverMigrate006)
}

func upRiverMigrate006(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 6,
	})
	return err
}

func downRiverMigrate006(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		TargetVersion: 5,
	})
	return err
}
