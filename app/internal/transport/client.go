package transport

import (
	"fmt"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client[C any] struct {
	client *C
	close  func() error
}

func (c *Client[C]) Client() *C {
	return c.client
}

func (c *Client[C]) Close() error {
	if c.close == nil {
		return nil
	}
	return c.close()
}

type GoaGRPCClientAdapter[C any] struct {
	newFunc GoaGRPCClientNewFunc[C]
}

type GoaGRPCClientNewFunc[C any] func(cc *grpc.ClientConn, opts ...grpc.CallOption) *C

func GoaGRPCClient[C any](newFunc GoaGRPCClientNewFunc[C]) *GoaGRPCClientAdapter[C] {
	return &GoaGRPCClientAdapter[C]{
		newFunc: newFunc,
	}
}

func (a *GoaGRPCClientAdapter[C]) Adapt(addr string) (*Client[C], error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("create gRPC client: %w", err)
	}

	return &Client[C]{
		close:  conn.Close,
		client: a.newFunc(conn),
	}, nil
}

type GoaHTTPClientAdapter[C any] struct {
	newFunc GoaHTTPClientNewFunc[C]
}

type GoaHTTPClientNewFunc[C any] func(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *C

func GoaHTTPClient[C any](newFunc GoaHTTPClientNewFunc[C]) *GoaHTTPClientAdapter[C] {
	return &GoaHTTPClientAdapter[C]{
		newFunc: newFunc,
	}
}

func (a *GoaHTTPClientAdapter[C]) Adapt(scheme, addr string) (*Client[C], error) {
	doer := http.DefaultClient
	enc := goahttp.RequestEncoder
	dec := goahttp.ResponseDecoder

	return &Client[C]{
		client: a.newFunc(scheme, addr, doer, enc, dec, false),
	}, nil
}
