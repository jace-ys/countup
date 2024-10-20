package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/jace-ys/countup/internal/postgres"
	"github.com/jace-ys/countup/schema/migrations"
)

type MigrateCmd struct {
	Command string   `arg:"" help:"Command to pass to goose migrate."`
	Args    []string `arg:"" optional:"" passthrough:"" help:"Additional args to pass to goose migrate."`

	Database struct {
		ConnectionURI string `env:"CONNECTION_URI" required:"" help:"Connection URI for connecting to the database."`
	} `embed:"" envprefix:"DATABASE_" prefix:"database."`

	Migrations struct {
		Dir     string `env:"DIR" default:"." help:"Path to the directory containing migration files."`
		LocalFS bool   `env:"LOCALFS" help:"Allows discovering of migration files from OS filesystem."`
	} `embed:"" envprefix:"MIGRATIONS_" prefix:"migrations."`
}

func (c *MigrateCmd) Run(ctx context.Context, g *Globals) error {
	db, err := postgres.NewPool(ctx, c.Database.ConnectionURI)
	if err != nil {
		return fmt.Errorf("init db pool: %w", err)
	}
	defer db.Close()

	conn := stdlib.OpenDBFromPool(db.Pool)
	defer conn.Close()

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}

	if !c.Migrations.LocalFS {
		goose.SetBaseFS(migrations.FSDir)
	}

	if err := goose.RunContext(ctx, c.Command, conn, c.Migrations.Dir, c.Args...); err != nil {
		return fmt.Errorf("run goose command: %w", err)
	}

	return nil
}
