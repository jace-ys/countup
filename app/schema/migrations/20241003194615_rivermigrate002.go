package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/riverqueue/river/rivermigrate"
)

func init() {
	goose.AddMigrationContext(upRiverMigrate002, downRiverMigrate002)
}

func upRiverMigrate002(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 2,
	})
	return err
}

func downRiverMigrate002(ctx context.Context, tx *sql.Tx) error {
	_, err := riverMigrate.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		TargetVersion: -1,
	})
	return err
}
