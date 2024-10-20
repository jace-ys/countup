package main

import (
	"io"
)

type Globals struct {
	Debug  bool      `env:"DEBUG" help:"Enable debug logging."`
	Writer io.Writer `kong:"-"`
}
