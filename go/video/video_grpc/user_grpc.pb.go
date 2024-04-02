// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user.proto

package video_grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_Users_FullMethodName              = "/user.UserService/Users"
	UserService_User_FullMethodName               = "/user.UserService/User"
	UserService_UserByAuth_FullMethodName         = "/user.UserService/UserByAuth"
	UserService_RegisterUser_FullMethodName       = "/user.UserService/RegisterUser"
	UserService_SubscribeChannel_FullMethodName   = "/user.UserService/SubscribeChannel"
	UserService_UnSubscribeChannel_FullMethodName = "/user.UserService/UnSubscribeChannel"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Users(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UsersResponse, error)
	User(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserPayload, error)
	UserByAuth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserPayload, error)
	RegisterUser(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*UserPayload, error)
	SubscribeChannel(ctx context.Context, in *SubscribeChannelInput, opts ...grpc.CallOption) (*SubscriptionPayload, error)
	UnSubscribeChannel(ctx context.Context, in *SubscribeChannelInput, opts ...grpc.CallOption) (*SubscriptionPayload, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Users(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, UserService_Users_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) User(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserPayload, error) {
	out := new(UserPayload)
	err := c.cc.Invoke(ctx, UserService_User_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserByAuth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserPayload, error) {
	out := new(UserPayload)
	err := c.cc.Invoke(ctx, UserService_UserByAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*UserPayload, error) {
	out := new(UserPayload)
	err := c.cc.Invoke(ctx, UserService_RegisterUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SubscribeChannel(ctx context.Context, in *SubscribeChannelInput, opts ...grpc.CallOption) (*SubscriptionPayload, error) {
	out := new(SubscriptionPayload)
	err := c.cc.Invoke(ctx, UserService_SubscribeChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnSubscribeChannel(ctx context.Context, in *SubscribeChannelInput, opts ...grpc.CallOption) (*SubscriptionPayload, error) {
	out := new(SubscriptionPayload)
	err := c.cc.Invoke(ctx, UserService_UnSubscribeChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Users(context.Context, *empty.Empty) (*UsersResponse, error)
	User(context.Context, *UserID) (*UserPayload, error)
	UserByAuth(context.Context, *empty.Empty) (*UserPayload, error)
	RegisterUser(context.Context, *UserInput) (*UserPayload, error)
	SubscribeChannel(context.Context, *SubscribeChannelInput) (*SubscriptionPayload, error)
	UnSubscribeChannel(context.Context, *SubscribeChannelInput) (*SubscriptionPayload, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Users(context.Context, *empty.Empty) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Users not implemented")
}
func (UnimplementedUserServiceServer) User(context.Context, *UserID) (*UserPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedUserServiceServer) UserByAuth(context.Context, *empty.Empty) (*UserPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserByAuth not implemented")
}
func (UnimplementedUserServiceServer) RegisterUser(context.Context, *UserInput) (*UserPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) SubscribeChannel(context.Context, *SubscribeChannelInput) (*SubscriptionPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeChannel not implemented")
}
func (UnimplementedUserServiceServer) UnSubscribeChannel(context.Context, *SubscribeChannelInput) (*SubscriptionPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribeChannel not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Users_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Users(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_User_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).User(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserByAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserByAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserByAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserByAuth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*UserInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SubscribeChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeChannelInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SubscribeChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SubscribeChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SubscribeChannel(ctx, req.(*SubscribeChannelInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnSubscribeChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeChannelInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnSubscribeChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UnSubscribeChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnSubscribeChannel(ctx, req.(*SubscribeChannelInput))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Users",
			Handler:    _UserService_Users_Handler,
		},
		{
			MethodName: "User",
			Handler:    _UserService_User_Handler,
		},
		{
			MethodName: "UserByAuth",
			Handler:    _UserService_UserByAuth_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "SubscribeChannel",
			Handler:    _UserService_SubscribeChannel_Handler,
		},
		{
			MethodName: "UnSubscribeChannel",
			Handler:    _UserService_UnSubscribeChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}