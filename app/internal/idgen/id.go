package idgen

import (
	"github.com/rs/xid"
)

func NewID(prefix string) string {
	id := xid.New()
	return prefix + "_" + id.String()
}
