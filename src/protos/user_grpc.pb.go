// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: src/protos/user.proto

package grpc_server1

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

// UserRPCClient is the client API for UserRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRPCClient interface {
	SelectUser(ctx context.Context, in *DataSelectUser, opts ...grpc.CallOption) (*ResponseSelectUser, error)
	SelectUsers(ctx context.Context, in *DataSelectUsers, opts ...grpc.CallOption) (*ResponseSelectUsers, error)
}

type userRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRPCClient(cc grpc.ClientConnInterface) UserRPCClient {
	return &userRPCClient{cc}
}

func (c *userRPCClient) SelectUser(ctx context.Context, in *DataSelectUser, opts ...grpc.CallOption) (*ResponseSelectUser, error) {
	out := new(ResponseSelectUser)
	err := c.cc.Invoke(ctx, "/protos.UserRPC/SelectUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) SelectUsers(ctx context.Context, in *DataSelectUsers, opts ...grpc.CallOption) (*ResponseSelectUsers, error) {
	out := new(ResponseSelectUsers)
	err := c.cc.Invoke(ctx, "/protos.UserRPC/SelectUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRPCServer is the server API for UserRPC service.
// All implementations must embed UnimplementedUserRPCServer
// for forward compatibility
type UserRPCServer interface {
	SelectUser(context.Context, *DataSelectUser) (*ResponseSelectUser, error)
	SelectUsers(context.Context, *DataSelectUsers) (*ResponseSelectUsers, error)
	mustEmbedUnimplementedUserRPCServer()
}

// UnimplementedUserRPCServer must be embedded to have forward compatible implementations.
type UnimplementedUserRPCServer struct {
}

func (UnimplementedUserRPCServer) SelectUser(context.Context, *DataSelectUser) (*ResponseSelectUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectUser not implemented")
}
func (UnimplementedUserRPCServer) SelectUsers(context.Context, *DataSelectUsers) (*ResponseSelectUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectUsers not implemented")
}
func (UnimplementedUserRPCServer) mustEmbedUnimplementedUserRPCServer() {}

// UnsafeUserRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRPCServer will
// result in compilation errors.
type UnsafeUserRPCServer interface {
	mustEmbedUnimplementedUserRPCServer()
}

func RegisterUserRPCServer(s grpc.ServiceRegistrar, srv UserRPCServer) {
	s.RegisterService(&UserRPC_ServiceDesc, srv)
}

func _UserRPC_SelectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSelectUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).SelectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserRPC/SelectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).SelectUser(ctx, req.(*DataSelectUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_SelectUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSelectUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).SelectUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserRPC/SelectUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).SelectUsers(ctx, req.(*DataSelectUsers))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRPC_ServiceDesc is the grpc.ServiceDesc for UserRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserRPC",
	HandlerType: (*UserRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelectUser",
			Handler:    _UserRPC_SelectUser_Handler,
		},
		{
			MethodName: "SelectUsers",
			Handler:    _UserRPC_SelectUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/protos/user.proto",
}
