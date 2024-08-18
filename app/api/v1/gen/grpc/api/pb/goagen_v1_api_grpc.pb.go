// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apipb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	// CounterGet implements CounterGet.
	CounterGet(ctx context.Context, in *CounterGetRequest, opts ...grpc.CallOption) (*CounterGetResponse, error)
	// CounterIncrement implements CounterIncrement.
	CounterIncrement(ctx context.Context, in *CounterIncrementRequest, opts ...grpc.CallOption) (*CounterIncrementResponse, error)
	// Echo implements Echo.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CounterGet(ctx context.Context, in *CounterGetRequest, opts ...grpc.CallOption) (*CounterGetResponse, error) {
	out := new(CounterGetResponse)
	err := c.cc.Invoke(ctx, "/api.API/CounterGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CounterIncrement(ctx context.Context, in *CounterIncrementRequest, opts ...grpc.CallOption) (*CounterIncrementResponse, error) {
	out := new(CounterIncrementResponse)
	err := c.cc.Invoke(ctx, "/api.API/CounterIncrement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/api.API/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	// CounterGet implements CounterGet.
	CounterGet(context.Context, *CounterGetRequest) (*CounterGetResponse, error)
	// CounterIncrement implements CounterIncrement.
	CounterIncrement(context.Context, *CounterIncrementRequest) (*CounterIncrementResponse, error)
	// Echo implements Echo.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) CounterGet(context.Context, *CounterGetRequest) (*CounterGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CounterGet not implemented")
}
func (UnimplementedAPIServer) CounterIncrement(context.Context, *CounterIncrementRequest) (*CounterIncrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CounterIncrement not implemented")
}
func (UnimplementedAPIServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_CounterGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CounterGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CounterGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/CounterGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CounterGet(ctx, req.(*CounterGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CounterIncrement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CounterIncrementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CounterIncrement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/CounterIncrement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CounterIncrement(ctx, req.(*CounterIncrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CounterGet",
			Handler:    _API_CounterGet_Handler,
		},
		{
			MethodName: "CounterIncrement",
			Handler:    _API_CounterIncrement_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _API_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_v1_api.proto",
}
