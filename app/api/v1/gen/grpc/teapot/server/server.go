// Code generated by goa v3.19.1, DO NOT EDIT.
//
// teapot gRPC server
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package server

import (
	"context"

	teapotpb "github.com/jace-ys/countup/api/v1/gen/grpc/teapot/pb"
	teapot "github.com/jace-ys/countup/api/v1/gen/teapot"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the teapotpb.TeapotServer interface.
type Server struct {
	EchoH goagrpc.UnaryHandler
	teapotpb.UnimplementedTeapotServer
}

// New instantiates the server struct with the teapot service endpoints.
func New(e *teapot.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		EchoH: NewEchoHandler(e.Echo, uh),
	}
}

// NewEchoHandler creates a gRPC handler which serves the "teapot" service
// "Echo" endpoint.
func NewEchoHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeEchoRequest, EncodeEchoResponse)
	}
	return h
}

// Echo implements the "Echo" method in teapotpb.TeapotServer interface.
func (s *Server) Echo(ctx context.Context, message *teapotpb.EchoRequest) (*teapotpb.EchoResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Echo")
	ctx = context.WithValue(ctx, goa.ServiceKey, "teapot")
	resp, err := s.EchoH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*teapotpb.EchoResponse), nil
}