package healthz

import (
	"context"
	"fmt"

	"github.com/alexliesenfeld/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func GRPCCheck(name, target string) health.Check {
	return health.Check{
		Name: "grpc:" + name,
		Check: func(ctx context.Context) error {
			opts := []grpc.DialOption{
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			}

			conn, err := grpc.NewClient(target, opts...)
			if err != nil {
				return fmt.Errorf("create gRPC client: %w", err)
			}
			defer conn.Close()

			res, err := healthpb.NewHealthClient(conn).Check(ctx, &healthpb.HealthCheckRequest{})
			if err != nil {
				return fmt.Errorf("send gRPC request: %w", err)
			}

			if res.GetStatus() != healthpb.HealthCheckResponse_SERVING {
				return fmt.Errorf("gRPC service reported as non-serving: %q", res.GetStatus().String())
			}

			return nil
		},
	}
}
