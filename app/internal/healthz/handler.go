package healthz

import (
	"context"
	"net/http"

	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func NewHTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

type GRPCHandler struct {
}

func NewGRPCHandler() healthpb.HealthServer {
	return &GRPCHandler{}
}

func (h *GRPCHandler) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	}, nil
}

func (h *GRPCHandler) Watch(req *healthpb.HealthCheckRequest, server healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
