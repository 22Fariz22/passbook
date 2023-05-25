// Code generated by MockGen. DO NOT EDIT.
// Source: server/proto/user_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	passbook "github.com/22Fariz22/passbook/server/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockUserServiceClient) AddAccount(ctx context.Context, in *passbook.AddAccountRequest, opts ...grpc.CallOption) (*passbook.AddAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAccount", varargs...)
	ret0, _ := ret[0].(*passbook.AddAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockUserServiceClientMockRecorder) AddAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockUserServiceClient)(nil).AddAccount), varargs...)
}

// AddBinary mocks base method.
func (m *MockUserServiceClient) AddBinary(ctx context.Context, in *passbook.AddBinaryRequest, opts ...grpc.CallOption) (*passbook.AddBinaryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddBinary", varargs...)
	ret0, _ := ret[0].(*passbook.AddBinaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBinary indicates an expected call of AddBinary.
func (mr *MockUserServiceClientMockRecorder) AddBinary(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinary", reflect.TypeOf((*MockUserServiceClient)(nil).AddBinary), varargs...)
}

// AddCard mocks base method.
func (m *MockUserServiceClient) AddCard(ctx context.Context, in *passbook.AddCardRequest, opts ...grpc.CallOption) (*passbook.AddCardResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCard", varargs...)
	ret0, _ := ret[0].(*passbook.AddCardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCard indicates an expected call of AddCard.
func (mr *MockUserServiceClientMockRecorder) AddCard(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockUserServiceClient)(nil).AddCard), varargs...)
}

// AddText mocks base method.
func (m *MockUserServiceClient) AddText(ctx context.Context, in *passbook.AddTextRequest, opts ...grpc.CallOption) (*passbook.AddTextResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddText", varargs...)
	ret0, _ := ret[0].(*passbook.AddTextResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddText indicates an expected call of AddText.
func (mr *MockUserServiceClientMockRecorder) AddText(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddText", reflect.TypeOf((*MockUserServiceClient)(nil).AddText), varargs...)
}

// FindByID mocks base method.
func (m *MockUserServiceClient) FindByID(ctx context.Context, in *passbook.FindByIDRequest, opts ...grpc.CallOption) (*passbook.FindByIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByID", varargs...)
	ret0, _ := ret[0].(*passbook.FindByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserServiceClientMockRecorder) FindByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserServiceClient)(nil).FindByID), varargs...)
}

// FindByLogin mocks base method.
func (m *MockUserServiceClient) FindByLogin(ctx context.Context, in *passbook.FindByLoginRequest, opts ...grpc.CallOption) (*passbook.FindByLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByLogin", varargs...)
	ret0, _ := ret[0].(*passbook.FindByLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *MockUserServiceClientMockRecorder) FindByLogin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockUserServiceClient)(nil).FindByLogin), varargs...)
}

// GetByTitle mocks base method.
func (m *MockUserServiceClient) GetByTitle(ctx context.Context, in *passbook.GetByTitleRequest, opts ...grpc.CallOption) (*passbook.GetByTitleResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByTitle", varargs...)
	ret0, _ := ret[0].(*passbook.GetByTitleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTitle indicates an expected call of GetByTitle.
func (mr *MockUserServiceClientMockRecorder) GetByTitle(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTitle", reflect.TypeOf((*MockUserServiceClient)(nil).GetByTitle), varargs...)
}

// GetFullList mocks base method.
func (m *MockUserServiceClient) GetFullList(ctx context.Context, in *passbook.GetFullListRequest, opts ...grpc.CallOption) (*passbook.GetFullListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFullList", varargs...)
	ret0, _ := ret[0].(*passbook.GetFullListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullList indicates an expected call of GetFullList.
func (mr *MockUserServiceClientMockRecorder) GetFullList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullList", reflect.TypeOf((*MockUserServiceClient)(nil).GetFullList), varargs...)
}

// GetMe mocks base method.
func (m *MockUserServiceClient) GetMe(ctx context.Context, in *passbook.GetMeRequest, opts ...grpc.CallOption) (*passbook.GetMeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMe", varargs...)
	ret0, _ := ret[0].(*passbook.GetMeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMe indicates an expected call of GetMe.
func (mr *MockUserServiceClientMockRecorder) GetMe(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockUserServiceClient)(nil).GetMe), varargs...)
}

// Login mocks base method.
func (m *MockUserServiceClient) Login(ctx context.Context, in *passbook.LoginRequest, opts ...grpc.CallOption) (*passbook.LoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*passbook.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceClientMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserServiceClient)(nil).Login), varargs...)
}

// Logout mocks base method.
func (m *MockUserServiceClient) Logout(ctx context.Context, in *passbook.LogoutRequest, opts ...grpc.CallOption) (*passbook.LogoutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Logout", varargs...)
	ret0, _ := ret[0].(*passbook.LogoutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout.
func (mr *MockUserServiceClientMockRecorder) Logout(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserServiceClient)(nil).Logout), varargs...)
}

// Register mocks base method.
func (m *MockUserServiceClient) Register(ctx context.Context, in *passbook.RegisterRequest, opts ...grpc.CallOption) (*passbook.RegisterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*passbook.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserServiceClient)(nil).Register), varargs...)
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockUserServiceServer) AddAccount(arg0 context.Context, arg1 *passbook.AddAccountRequest) (*passbook.AddAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccount", arg0, arg1)
	ret0, _ := ret[0].(*passbook.AddAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockUserServiceServerMockRecorder) AddAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockUserServiceServer)(nil).AddAccount), arg0, arg1)
}

// AddBinary mocks base method.
func (m *MockUserServiceServer) AddBinary(arg0 context.Context, arg1 *passbook.AddBinaryRequest) (*passbook.AddBinaryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinary", arg0, arg1)
	ret0, _ := ret[0].(*passbook.AddBinaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBinary indicates an expected call of AddBinary.
func (mr *MockUserServiceServerMockRecorder) AddBinary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinary", reflect.TypeOf((*MockUserServiceServer)(nil).AddBinary), arg0, arg1)
}

// AddCard mocks base method.
func (m *MockUserServiceServer) AddCard(arg0 context.Context, arg1 *passbook.AddCardRequest) (*passbook.AddCardResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCard", arg0, arg1)
	ret0, _ := ret[0].(*passbook.AddCardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCard indicates an expected call of AddCard.
func (mr *MockUserServiceServerMockRecorder) AddCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockUserServiceServer)(nil).AddCard), arg0, arg1)
}

// AddText mocks base method.
func (m *MockUserServiceServer) AddText(arg0 context.Context, arg1 *passbook.AddTextRequest) (*passbook.AddTextResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddText", arg0, arg1)
	ret0, _ := ret[0].(*passbook.AddTextResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddText indicates an expected call of AddText.
func (mr *MockUserServiceServerMockRecorder) AddText(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddText", reflect.TypeOf((*MockUserServiceServer)(nil).AddText), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockUserServiceServer) FindByID(arg0 context.Context, arg1 *passbook.FindByIDRequest) (*passbook.FindByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*passbook.FindByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserServiceServerMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserServiceServer)(nil).FindByID), arg0, arg1)
}

// FindByLogin mocks base method.
func (m *MockUserServiceServer) FindByLogin(arg0 context.Context, arg1 *passbook.FindByLoginRequest) (*passbook.FindByLoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", arg0, arg1)
	ret0, _ := ret[0].(*passbook.FindByLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *MockUserServiceServerMockRecorder) FindByLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockUserServiceServer)(nil).FindByLogin), arg0, arg1)
}

// GetByTitle mocks base method.
func (m *MockUserServiceServer) GetByTitle(arg0 context.Context, arg1 *passbook.GetByTitleRequest) (*passbook.GetByTitleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTitle", arg0, arg1)
	ret0, _ := ret[0].(*passbook.GetByTitleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTitle indicates an expected call of GetByTitle.
func (mr *MockUserServiceServerMockRecorder) GetByTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTitle", reflect.TypeOf((*MockUserServiceServer)(nil).GetByTitle), arg0, arg1)
}

// GetFullList mocks base method.
func (m *MockUserServiceServer) GetFullList(arg0 context.Context, arg1 *passbook.GetFullListRequest) (*passbook.GetFullListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullList", arg0, arg1)
	ret0, _ := ret[0].(*passbook.GetFullListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullList indicates an expected call of GetFullList.
func (mr *MockUserServiceServerMockRecorder) GetFullList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullList", reflect.TypeOf((*MockUserServiceServer)(nil).GetFullList), arg0, arg1)
}

// GetMe mocks base method.
func (m *MockUserServiceServer) GetMe(arg0 context.Context, arg1 *passbook.GetMeRequest) (*passbook.GetMeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMe", arg0, arg1)
	ret0, _ := ret[0].(*passbook.GetMeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMe indicates an expected call of GetMe.
func (mr *MockUserServiceServerMockRecorder) GetMe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockUserServiceServer)(nil).GetMe), arg0, arg1)
}

// Login mocks base method.
func (m *MockUserServiceServer) Login(arg0 context.Context, arg1 *passbook.LoginRequest) (*passbook.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*passbook.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceServerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserServiceServer)(nil).Login), arg0, arg1)
}

// Logout mocks base method.
func (m *MockUserServiceServer) Logout(arg0 context.Context, arg1 *passbook.LogoutRequest) (*passbook.LogoutResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0, arg1)
	ret0, _ := ret[0].(*passbook.LogoutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout.
func (mr *MockUserServiceServerMockRecorder) Logout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserServiceServer)(nil).Logout), arg0, arg1)
}

// Register mocks base method.
func (m *MockUserServiceServer) Register(arg0 context.Context, arg1 *passbook.RegisterRequest) (*passbook.RegisterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*passbook.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserServiceServer)(nil).Register), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}