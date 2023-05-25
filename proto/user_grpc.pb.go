// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/user.proto

package passbook

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
	UserService_Register_FullMethodName    = "/passbook.UserService/Register"
	UserService_FindByLogin_FullMethodName = "/passbook.UserService/FindByLogin"
	UserService_FindByID_FullMethodName    = "/passbook.UserService/FindByID"
	UserService_Login_FullMethodName       = "/passbook.UserService/Login"
	UserService_GetMe_FullMethodName       = "/passbook.UserService/GetMe"
	UserService_Logout_FullMethodName      = "/passbook.UserService/Logout"
	UserService_AddAccount_FullMethodName  = "/passbook.UserService/AddAccount"
	UserService_AddText_FullMethodName     = "/passbook.UserService/AddText"
	UserService_AddBinary_FullMethodName   = "/passbook.UserService/AddBinary"
	UserService_AddCard_FullMethodName     = "/passbook.UserService/AddCard"
	UserService_GetByTitle_FullMethodName  = "/passbook.UserService/GetByTitle"
	UserService_GetFullList_FullMethodName = "/passbook.UserService/GetFullList"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	FindByLogin(ctx context.Context, in *FindByLoginRequest, opts ...grpc.CallOption) (*FindByLoginResponse, error)
	FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*GetMeResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AddAccountResponse, error)
	AddText(ctx context.Context, in *AddTextRequest, opts ...grpc.CallOption) (*AddTextResponse, error)
	AddBinary(ctx context.Context, in *AddBinaryRequest, opts ...grpc.CallOption) (*AddBinaryResponse, error)
	AddCard(ctx context.Context, in *AddCardRequest, opts ...grpc.CallOption) (*AddCardResponse, error)
	GetByTitle(ctx context.Context, in *GetByTitleRequest, opts ...grpc.CallOption) (*GetByTitleResponse, error)
	GetFullList(ctx context.Context, in *GetFullListRequest, opts ...grpc.CallOption) (*GetFullListResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindByLogin(ctx context.Context, in *FindByLoginRequest, opts ...grpc.CallOption) (*FindByLoginResponse, error) {
	out := new(FindByLoginResponse)
	err := c.cc.Invoke(ctx, UserService_FindByLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error) {
	out := new(FindByIDResponse)
	err := c.cc.Invoke(ctx, UserService_FindByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*GetMeResponse, error) {
	out := new(GetMeResponse)
	err := c.cc.Invoke(ctx, UserService_GetMe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, UserService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AddAccountResponse, error) {
	out := new(AddAccountResponse)
	err := c.cc.Invoke(ctx, UserService_AddAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddText(ctx context.Context, in *AddTextRequest, opts ...grpc.CallOption) (*AddTextResponse, error) {
	out := new(AddTextResponse)
	err := c.cc.Invoke(ctx, UserService_AddText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddBinary(ctx context.Context, in *AddBinaryRequest, opts ...grpc.CallOption) (*AddBinaryResponse, error) {
	out := new(AddBinaryResponse)
	err := c.cc.Invoke(ctx, UserService_AddBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddCard(ctx context.Context, in *AddCardRequest, opts ...grpc.CallOption) (*AddCardResponse, error) {
	out := new(AddCardResponse)
	err := c.cc.Invoke(ctx, UserService_AddCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByTitle(ctx context.Context, in *GetByTitleRequest, opts ...grpc.CallOption) (*GetByTitleResponse, error) {
	out := new(GetByTitleResponse)
	err := c.cc.Invoke(ctx, UserService_GetByTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFullList(ctx context.Context, in *GetFullListRequest, opts ...grpc.CallOption) (*GetFullListResponse, error) {
	out := new(GetFullListResponse)
	err := c.cc.Invoke(ctx, UserService_GetFullList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	FindByLogin(context.Context, *FindByLoginRequest) (*FindByLoginResponse, error)
	FindByID(context.Context, *FindByIDRequest) (*FindByIDResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetMe(context.Context, *GetMeRequest) (*GetMeResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	AddAccount(context.Context, *AddAccountRequest) (*AddAccountResponse, error)
	AddText(context.Context, *AddTextRequest) (*AddTextResponse, error)
	AddBinary(context.Context, *AddBinaryRequest) (*AddBinaryResponse, error)
	AddCard(context.Context, *AddCardRequest) (*AddCardResponse, error)
	GetByTitle(context.Context, *GetByTitleRequest) (*GetByTitleResponse, error)
	GetFullList(context.Context, *GetFullListRequest) (*GetFullListResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) FindByLogin(context.Context, *FindByLoginRequest) (*FindByLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByLogin not implemented")
}
func (UnimplementedUserServiceServer) FindByID(context.Context, *FindByIDRequest) (*FindByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GetMe(context.Context, *GetMeRequest) (*GetMeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedUserServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServiceServer) AddAccount(context.Context, *AddAccountRequest) (*AddAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (UnimplementedUserServiceServer) AddText(context.Context, *AddTextRequest) (*AddTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddText not implemented")
}
func (UnimplementedUserServiceServer) AddBinary(context.Context, *AddBinaryRequest) (*AddBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBinary not implemented")
}
func (UnimplementedUserServiceServer) AddCard(context.Context, *AddCardRequest) (*AddCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCard not implemented")
}
func (UnimplementedUserServiceServer) GetByTitle(context.Context, *GetByTitleRequest) (*GetByTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByTitle not implemented")
}
func (UnimplementedUserServiceServer) GetFullList(context.Context, *GetFullListRequest) (*GetFullListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullList not implemented")
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

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindByLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindByLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindByLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindByLogin(ctx, req.(*FindByLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindByID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMe(ctx, req.(*GetMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAccount(ctx, req.(*AddAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddText(ctx, req.(*AddTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddBinary(ctx, req.(*AddBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddCard(ctx, req.(*AddCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByTitle(ctx, req.(*GetByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFullList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetFullList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetFullList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetFullList(ctx, req.(*GetFullListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "passbook.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "FindByLogin",
			Handler:    _UserService_FindByLogin_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _UserService_FindByID_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetMe",
			Handler:    _UserService_GetMe_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _UserService_AddAccount_Handler,
		},
		{
			MethodName: "AddText",
			Handler:    _UserService_AddText_Handler,
		},
		{
			MethodName: "AddBinary",
			Handler:    _UserService_AddBinary_Handler,
		},
		{
			MethodName: "AddCard",
			Handler:    _UserService_AddCard_Handler,
		},
		{
			MethodName: "GetByTitle",
			Handler:    _UserService_GetByTitle_Handler,
		},
		{
			MethodName: "GetFullList",
			Handler:    _UserService_GetFullList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}