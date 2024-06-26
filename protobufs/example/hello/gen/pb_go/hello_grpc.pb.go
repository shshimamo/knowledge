// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: hello.proto

package pb_go

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

const (
	HelloService_GetServerResponse_FullMethodName = "/hello.helloService/GetServerResponse"
	HelloService_Optional_FullMethodName          = "/hello.helloService/Optional"
	HelloService_OptionalOneOf_FullMethodName     = "/hello.helloService/OptionalOneOf"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	GetServerResponse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error)
	Optional(ctx context.Context, in *Message, opts ...grpc.CallOption) (*OptionalResponse, error)
	OptionalOneOf(ctx context.Context, in *Message, opts ...grpc.CallOption) (*OptionalOneOfResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) GetServerResponse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, HelloService_GetServerResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Optional(ctx context.Context, in *Message, opts ...grpc.CallOption) (*OptionalResponse, error) {
	out := new(OptionalResponse)
	err := c.cc.Invoke(ctx, HelloService_Optional_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) OptionalOneOf(ctx context.Context, in *Message, opts ...grpc.CallOption) (*OptionalOneOfResponse, error) {
	out := new(OptionalOneOfResponse)
	err := c.cc.Invoke(ctx, HelloService_OptionalOneOf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	GetServerResponse(context.Context, *Message) (*MessageResponse, error)
	Optional(context.Context, *Message) (*OptionalResponse, error)
	OptionalOneOf(context.Context, *Message) (*OptionalOneOfResponse, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) GetServerResponse(context.Context, *Message) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerResponse not implemented")
}
func (UnimplementedHelloServiceServer) Optional(context.Context, *Message) (*OptionalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Optional not implemented")
}
func (UnimplementedHelloServiceServer) OptionalOneOf(context.Context, *Message) (*OptionalOneOfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OptionalOneOf not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_GetServerResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).GetServerResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_GetServerResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).GetServerResponse(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Optional_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Optional(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_Optional_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Optional(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_OptionalOneOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).OptionalOneOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_OptionalOneOf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).OptionalOneOf(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.helloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerResponse",
			Handler:    _HelloService_GetServerResponse_Handler,
		},
		{
			MethodName: "Optional",
			Handler:    _HelloService_Optional_Handler,
		},
		{
			MethodName: "OptionalOneOf",
			Handler:    _HelloService_OptionalOneOf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
