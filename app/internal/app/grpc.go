package app

import (
	"context"
	"fmt"
	"net"

	"github.com/alexliesenfeld/health"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/reflection/grpc_reflection_v1"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/stats"

	"github.com/jace-ys/countup/internal/healthz"
	"github.com/jace-ys/countup/internal/slog"
	"github.com/jace-ys/countup/internal/transport/middleware/idgen"
	"github.com/jace-ys/countup/internal/transport/middleware/recovery"
)

type GRPCServer[SS any] struct {
	name string
	addr string
	srv  *grpc.Server
}

func NewGRPCServer[SS any](ctx context.Context, name string, port int) *GRPCServer[SS] {
	addr := fmt.Sprintf(":%d", port)

	excludedMethods := map[string]bool{
		grpc_reflection_v1.ServerReflection_ServerReflectionInfo_FullMethodName:      true,
		grpc_reflection_v1alpha.ServerReflection_ServerReflectionInfo_FullMethodName: true,
		healthpb.Health_Check_FullMethodName:                                         true,
		healthpb.Health_Watch_FullMethodName:                                         true,
	}

	logCtx := log.With(ctx, slog.KV("server", name))
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(logCtx),
			withMethodFilter(idgen.UnaryServerInterceptor(), excludedMethods),
			withMethodFilter(slog.UnaryServerInterceptor(logCtx), excludedMethods),
			withMethodFilter(debug.UnaryServerInterceptor(), excludedMethods),
		),
		grpc.StatsHandler(otelgrpc.NewServerHandler(
			otelgrpc.WithSpanAttributes(attribute.String("rpc.server.name", name)),
			otelgrpc.WithFilter(func(info *stats.RPCTagInfo) bool {
				return !excludedMethods[info.FullMethodName]
			}),
		)),
	)

	reflection.Register(srv)
	healthpb.RegisterHealthServer(srv, healthz.NewGRPCHandler())

	return &GRPCServer[SS]{
		name: name,
		addr: addr,
		srv:  srv,
	}
}

func withMethodFilter(interceptor grpc.UnaryServerInterceptor, excluded map[string]bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if excluded := excluded[info.FullMethod]; excluded {
			return handler(ctx, req)
		}
		return interceptor(ctx, req, info, handler)
	}
}

func (s *GRPCServer[SS]) RegisterHandler(sd *grpc.ServiceDesc, ss SS) {
	s.srv.RegisterService(sd, ss)
}

var _ Server = (*GRPCServer[any])(nil)

func (s *GRPCServer[SS]) Name() string {
	return s.name
}

func (s *GRPCServer[SS]) Kind() string {
	return "grpc"
}

func (s *GRPCServer[SS]) Addr() string {
	return s.addr
}

func (s *GRPCServer[SS]) Serve(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("tcp listener: %w", err)
	}

	if err := s.srv.Serve(lis); err != nil {
		return fmt.Errorf("serving grpc server: %w", err)
	}

	return nil
}

func (s *GRPCServer[SS]) Shutdown(ctx context.Context) error {
	ok := make(chan struct{})

	go func() {
		s.srv.GracefulStop()
		close(ok)
	}()

	select {
	case <-ok:
		return nil
	case <-ctx.Done():
		s.srv.Stop()
		return ctx.Err() //nolint:wrapcheck
	}
}

var _ healthz.Target = (*GRPCServer[any])(nil)

func (s *GRPCServer[SS]) HealthChecks() []health.Check {
	return []health.Check{
		healthz.GRPCCheck(s.Name(), s.Addr()),
	}
}
