package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/slog"
)

type RootCmd struct {
	Globals

	Migrate MigrateCmd `cmd:"" help:"Run database migrations."`
	Server  ServerCmd  `cmd:"" help:"Run the countup server."`
	Version VersionCmd `cmd:"" help:"Show version information."`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ctx.Done()
		stop()
	}()

	var root RootCmd
	cli := kong.Parse(&root,
		kong.Name(api.APIName),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:             true,
			FlagsLast:           true,
			NoExpandSubcommands: true,
		}),
	)

	root.Writer = os.Stdout
	ctx = slog.NewContext(ctx, root.Writer, root.Debug)

	cli.BindTo(ctx, (*context.Context)(nil))
	cli.FatalIfErrorf(cli.Run(&root.Globals))
}
