// Code generated by goa v3.19.1, DO NOT EDIT.
//
// api gRPC server encoders and decoders
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package server

import (
	"context"

	api "github.com/jace-ys/countup/api/v1/gen/api"
	apiviews "github.com/jace-ys/countup/api/v1/gen/api/views"
	apipb "github.com/jace-ys/countup/api/v1/gen/grpc/api/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeCounterGetResponse encodes responses from the "api" service
// "CounterGet" endpoint.
func EncodeCounterGetResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*apiviews.CounterInfo)
	if !ok {
		return nil, goagrpc.ErrInvalidType("api", "CounterGet", "*apiviews.CounterInfo", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoCounterGetResponse(result)
	return resp, nil
}

// EncodeCounterIncrementResponse encodes responses from the "api" service
// "CounterIncrement" endpoint.
func EncodeCounterIncrementResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*apiviews.CounterInfo)
	if !ok {
		return nil, goagrpc.ErrInvalidType("api", "CounterIncrement", "*apiviews.CounterInfo", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoCounterIncrementResponse(result)
	return resp, nil
}

// DecodeCounterIncrementRequest decodes requests sent to "api" service
// "CounterIncrement" endpoint.
func DecodeCounterIncrementRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *apipb.CounterIncrementRequest
		ok      bool
	)
	{
		if message, ok = v.(*apipb.CounterIncrementRequest); !ok {
			return nil, goagrpc.ErrInvalidType("api", "CounterIncrement", "*apipb.CounterIncrementRequest", v)
		}
	}
	var payload *api.CounterIncrementPayload
	{
		payload = NewCounterIncrementPayload(message)
	}
	return payload, nil
}

// EncodeEchoResponse encodes responses from the "api" service "Echo" endpoint.
func EncodeEchoResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*api.EchoResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("api", "Echo", "*api.EchoResult", v)
	}
	resp := NewProtoEchoResponse(result)
	return resp, nil
}

// DecodeEchoRequest decodes requests sent to "api" service "Echo" endpoint.
func DecodeEchoRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *apipb.EchoRequest
		ok      bool
	)
	{
		if message, ok = v.(*apipb.EchoRequest); !ok {
			return nil, goagrpc.ErrInvalidType("api", "Echo", "*apipb.EchoRequest", v)
		}
	}
	var payload *api.EchoPayload
	{
		payload = NewEchoPayload(message)
	}
	return payload, nil
}
