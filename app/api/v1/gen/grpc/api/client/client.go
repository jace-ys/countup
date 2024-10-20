// Code generated by goa v3.19.1, DO NOT EDIT.
//
// api gRPC client
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package client

import (
	"context"

	apipb "github.com/jace-ys/countup/api/v1/gen/grpc/api/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli apipb.APIClient
	opts    []grpc.CallOption
} // NewClient instantiates gRPC client for all the api service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: apipb.NewAPIClient(cc),
		opts:    opts,
	}
} // AuthToken calls the "AuthToken" function in apipb.APIClient interface.
func (c *Client) AuthToken() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildAuthTokenFunc(c.grpccli, c.opts...),
			EncodeAuthTokenRequest,
			DecodeAuthTokenResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
} // CounterGet calls the "CounterGet" function in apipb.APIClient interface.
func (c *Client) CounterGet() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildCounterGetFunc(c.grpccli, c.opts...),
			nil,
			DecodeCounterGetResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
} // CounterIncrement calls the "CounterIncrement" function in apipb.APIClient
// interface.
func (c *Client) CounterIncrement() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildCounterIncrementFunc(c.grpccli, c.opts...),
			EncodeCounterIncrementRequest,
			DecodeCounterIncrementResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
