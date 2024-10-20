package migrations

import (
	"embed"
)

//go:embed *.sql
var FSDir embed.FS
