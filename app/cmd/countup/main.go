package main

import (
	"context"
	"os"

	"github.com/alecthomas/kong"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/ctxlog"
)

type RootCmd struct {
	Globals

	Migrate MigrateCmd `cmd:"" help:"Run database migrations."`
	Server  ServerCmd  `cmd:"" help:"Run the countup server."`
	Version VersionCmd `cmd:"" help:"Show version information."`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
	ctx = ctxlog.NewContext(ctx, root.Writer, root.Debug)

	cli.BindTo(ctx, (*context.Context)(nil))
	cli.FatalIfErrorf(cli.Run(&root.Globals))
}
