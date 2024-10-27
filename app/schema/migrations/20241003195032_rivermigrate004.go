package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/riverqueue/river/rivermigrate"
)

func init() {
	goose.AddMigrationContext(upRiverMigrate004, downRiverMigrate004)
}

func upRiverMigrate004(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 4,
	})
	return err
}

func downRiverMigrate004(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		TargetVersion: 3,
	})
	return err
}
