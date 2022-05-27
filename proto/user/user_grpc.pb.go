// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user.proto

package user

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdReply, error)
	GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserReply, error)
	GetUsersByApplicationId(ctx context.Context, in *GetUsersByApplicationIdRequest, opts ...grpc.CallOption) (*GetUsersByApplicationIdReply, error)
	GetAvailableUsers(ctx context.Context, in *GetAvailableUsersRequest, opts ...grpc.CallOption) (*GetAvailableUsersReply, error)
	GetAvailableUsersByApplicationId(ctx context.Context, in *GetAvailableUsersByApplicationIdRequest, opts ...grpc.CallOption) (*GetAvailableUsersByApplicationIdReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdReply, error) {
	out := new(GetUserByIdReply)
	err := c.cc.Invoke(ctx, "/user.User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserReply, error) {
	out := new(GetAllUserReply)
	err := c.cc.Invoke(ctx, "/user.User/GetAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsersByApplicationId(ctx context.Context, in *GetUsersByApplicationIdRequest, opts ...grpc.CallOption) (*GetUsersByApplicationIdReply, error) {
	out := new(GetUsersByApplicationIdReply)
	err := c.cc.Invoke(ctx, "/user.User/GetUsersByApplicationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAvailableUsers(ctx context.Context, in *GetAvailableUsersRequest, opts ...grpc.CallOption) (*GetAvailableUsersReply, error) {
	out := new(GetAvailableUsersReply)
	err := c.cc.Invoke(ctx, "/user.User/GetAvailableUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAvailableUsersByApplicationId(ctx context.Context, in *GetAvailableUsersByApplicationIdRequest, opts ...grpc.CallOption) (*GetAvailableUsersByApplicationIdReply, error) {
	out := new(GetAvailableUsersByApplicationIdReply)
	err := c.cc.Invoke(ctx, "/user.User/GetAvailableUsersByApplicationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdReply, error)
	GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserReply, error)
	GetUsersByApplicationId(context.Context, *GetUsersByApplicationIdRequest) (*GetUsersByApplicationIdReply, error)
	GetAvailableUsers(context.Context, *GetAvailableUsersRequest) (*GetAvailableUsersReply, error)
	GetAvailableUsersByApplicationId(context.Context, *GetAvailableUsersByApplicationIdRequest) (*GetAvailableUsersByApplicationIdReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServer) GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
}
func (UnimplementedUserServer) GetUsersByApplicationId(context.Context, *GetUsersByApplicationIdRequest) (*GetUsersByApplicationIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByApplicationId not implemented")
}
func (UnimplementedUserServer) GetAvailableUsers(context.Context, *GetAvailableUsersRequest) (*GetAvailableUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableUsers not implemented")
}
func (UnimplementedUserServer) GetAvailableUsersByApplicationId(context.Context, *GetAvailableUsersByApplicationIdRequest) (*GetAvailableUsersByApplicationIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableUsersByApplicationId not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAllUser(ctx, req.(*GetAllUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsersByApplicationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByApplicationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsersByApplicationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUsersByApplicationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsersByApplicationId(ctx, req.(*GetUsersByApplicationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAvailableUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAvailableUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetAvailableUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAvailableUsers(ctx, req.(*GetAvailableUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAvailableUsersByApplicationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableUsersByApplicationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAvailableUsersByApplicationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetAvailableUsersByApplicationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAvailableUsersByApplicationId(ctx, req.(*GetAvailableUsersByApplicationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "GetAllUser",
			Handler:    _User_GetAllUser_Handler,
		},
		{
			MethodName: "GetUsersByApplicationId",
			Handler:    _User_GetUsersByApplicationId_Handler,
		},
		{
			MethodName: "GetAvailableUsers",
			Handler:    _User_GetAvailableUsers_Handler,
		},
		{
			MethodName: "GetAvailableUsersByApplicationId",
			Handler:    _User_GetAvailableUsersByApplicationId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
