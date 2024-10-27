package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/riverqueue/river/rivermigrate"
)

func init() {
	goose.AddMigrationContext(upRiverMigrate005, downRiverMigrate005)
}

func upRiverMigrate005(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 5,
	})
	return err
}

func downRiverMigrate005(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		TargetVersion: 4,
	})
	return err
}
