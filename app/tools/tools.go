//go:build tools

package tools

import (
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
	_ "goa.design/goa/v3/cmd/goa"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
