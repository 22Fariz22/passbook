// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/22Fariz22/passbook/server/internal/entity"
	passbook "github.com/22Fariz22/passbook/server/proto"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockUserUseCase) AddAccount(ctx context.Context, userID string, request *passbook.AddAccountRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccount", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockUserUseCaseMockRecorder) AddAccount(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockUserUseCase)(nil).AddAccount), ctx, userID, request)
}

// AddBinary mocks base method.
func (m *MockUserUseCase) AddBinary(ctx context.Context, userID string, request *passbook.AddBinaryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinary", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinary indicates an expected call of AddBinary.
func (mr *MockUserUseCaseMockRecorder) AddBinary(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinary", reflect.TypeOf((*MockUserUseCase)(nil).AddBinary), ctx, userID, request)
}

// AddCard mocks base method.
func (m *MockUserUseCase) AddCard(ctx context.Context, userID string, request *passbook.AddCardRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCard", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCard indicates an expected call of AddCard.
func (mr *MockUserUseCaseMockRecorder) AddCard(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockUserUseCase)(nil).AddCard), ctx, userID, request)
}

// AddText mocks base method.
func (m *MockUserUseCase) AddText(ctx context.Context, userID string, request *passbook.AddTextRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddText", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddText indicates an expected call of AddText.
func (mr *MockUserUseCaseMockRecorder) AddText(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddText", reflect.TypeOf((*MockUserUseCase)(nil).AddText), ctx, userID, request)
}

// FindByID mocks base method.
func (m *MockUserUseCase) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, userID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserUseCaseMockRecorder) FindByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserUseCase)(nil).FindByID), ctx, userID)
}

// FindByLogin mocks base method.
func (m *MockUserUseCase) FindByLogin(ctx context.Context, login string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", ctx, login)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *MockUserUseCaseMockRecorder) FindByLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockUserUseCase)(nil).FindByLogin), ctx, login)
}

// GetByTitle mocks base method.
func (m *MockUserUseCase) GetByTitle(ctx context.Context, userID string, request *passbook.GetByTitleRequest) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTitle", ctx, userID, request)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTitle indicates an expected call of GetByTitle.
func (mr *MockUserUseCaseMockRecorder) GetByTitle(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTitle", reflect.TypeOf((*MockUserUseCase)(nil).GetByTitle), ctx, userID, request)
}

// GetFullList mocks base method.
func (m *MockUserUseCase) GetFullList(ctx context.Context, userID string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullList", ctx, userID)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullList indicates an expected call of GetFullList.
func (mr *MockUserUseCaseMockRecorder) GetFullList(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullList", reflect.TypeOf((*MockUserUseCase)(nil).GetFullList), ctx, userID)
}

// Login mocks base method.
func (m *MockUserUseCase) Login(ctx context.Context, email, password string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUseCaseMockRecorder) Login(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUseCase)(nil).Login), ctx, email, password)
}

// Register mocks base method.
func (m *MockUserUseCase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserUseCaseMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserUseCase)(nil).Register), ctx, user)
}
