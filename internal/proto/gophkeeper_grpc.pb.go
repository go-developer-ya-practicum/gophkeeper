// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: gophkeeper.proto

package proto

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

// GophKeeperClient is the client API for GophKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophKeeperClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
}

type gophKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperClient(cc grpc.ClientConnInterface) GophKeeperClient {
	return &gophKeeperClient{cc}
}

func (c *gophKeeperClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/proto.GophKeeper/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/proto.GophKeeper/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.GophKeeper/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophKeeperServer is the server API for GophKeeper service.
// All implementations must embed UnimplementedGophKeeperServer
// for forward compatibility
type GophKeeperServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
	mustEmbedUnimplementedGophKeeperServer()
}

// UnimplementedGophKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophKeeperServer struct {
}

func (UnimplementedGophKeeperServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedGophKeeperServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedGophKeeperServer) VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedGophKeeperServer) mustEmbedUnimplementedGophKeeperServer() {}

// UnsafeGophKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperServer will
// result in compilation errors.
type UnsafeGophKeeperServer interface {
	mustEmbedUnimplementedGophKeeperServer()
}

func RegisterGophKeeperServer(s grpc.ServiceRegistrar, srv GophKeeperServer) {
	s.RegisterService(&GophKeeper_ServiceDesc, srv)
}

func _GophKeeper_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GophKeeper/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GophKeeper/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GophKeeper/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).VerifyToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GophKeeper_ServiceDesc is the grpc.ServiceDesc for GophKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GophKeeper",
	HandlerType: (*GophKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _GophKeeper_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _GophKeeper_SignIn_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _GophKeeper_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gophkeeper.proto",
}
