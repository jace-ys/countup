package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/jace-ys/countup/internal/versioninfo"
)

type VersionCmd struct {
}

func (c *VersionCmd) Run(ctx context.Context, g *Globals) error {
	fmt.Printf(
		"Version: %v\nCommit SHA: %v\nGo Version: %v\nGo OS/Arch: %v/%v\n",
		versioninfo.Version, versioninfo.CommitSHA, runtime.Version(), runtime.GOOS, runtime.GOARCH,
	)
	return nil
}
