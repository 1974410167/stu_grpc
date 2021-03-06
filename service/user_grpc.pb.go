// Code generated by protoc-gen-go-s_grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-s_grpc v1.2.0
// - protoc             v3.6.1
// source: user.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the s_grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServerClient is the client API for UserServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServerClient interface {
	// 定义方法
	GetUserMessage(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	AddUserMessage(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type userServerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServerClient(cc grpc.ClientConnInterface) UserServerClient {
	return &userServerClient{cc}
}

func (c *userServerClient) GetUserMessage(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/service.UserServer/GetUserMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) AddUserMessage(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/service.UserServer/AddUserMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServer is the server API for UserServer service.
// All implementations must embed UnimplementedUserServerServer
// for forward compatibility
type UserServerServer interface {
	// 定义方法
	GetUserMessage(context.Context, *UserRequest) (*User, error)
	AddUserMessage(context.Context, *User) (*User, error)
	mustEmbedUnimplementedUserServerServer()
}

// UnimplementedUserServerServer must be embedded to have forward compatible implementations.
type UnimplementedUserServerServer struct {
}

func (UnimplementedUserServerServer) GetUserMessage(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMessage not implemented")
}
func (UnimplementedUserServerServer) AddUserMessage(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserMessage not implemented")
}
func (UnimplementedUserServerServer) mustEmbedUnimplementedUserServerServer() {}

// UnsafeUserServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServerServer will
// result in compilation errors.
type UnsafeUserServerServer interface {
	mustEmbedUnimplementedUserServerServer()
}

func RegisterUserServerServer(s grpc.ServiceRegistrar, srv UserServerServer) {
	s.RegisterService(&UserServer_ServiceDesc, srv)
}

func _UserServer_GetUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).GetUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserServer/GetUserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).GetUserMessage(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_AddUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).AddUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserServer/AddUserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).AddUserMessage(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServer_ServiceDesc is the grpc.ServiceDesc for UserServer service.
// It's only intended for direct use with s_grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.UserServer",
	HandlerType: (*UserServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserMessage",
			Handler:    _UserServer_GetUserMessage_Handler,
		},
		{
			MethodName: "AddUserMessage",
			Handler:    _UserServer_AddUserMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
