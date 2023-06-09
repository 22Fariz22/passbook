// Code generated by MockGen. DO NOT EDIT.
// Source: pg-repository.go

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

// MockUserPGRepository is a mock of UserPGRepository interface.
type MockUserPGRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserPGRepositoryMockRecorder
}

// MockUserPGRepositoryMockRecorder is the mock recorder for MockUserPGRepository.
type MockUserPGRepositoryMockRecorder struct {
	mock *MockUserPGRepository
}

// NewMockUserPGRepository creates a new mock instance.
func NewMockUserPGRepository(ctrl *gomock.Controller) *MockUserPGRepository {
	mock := &MockUserPGRepository{ctrl: ctrl}
	mock.recorder = &MockUserPGRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserPGRepository) EXPECT() *MockUserPGRepositoryMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockUserPGRepository) AddAccount(ctx context.Context, userID string, request *passbook.AddAccountRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccount", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockUserPGRepositoryMockRecorder) AddAccount(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockUserPGRepository)(nil).AddAccount), ctx, userID, request)
}

// AddBinary mocks base method.
func (m *MockUserPGRepository) AddBinary(ctx context.Context, userID string, request *passbook.AddBinaryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinary", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinary indicates an expected call of AddBinary.
func (mr *MockUserPGRepositoryMockRecorder) AddBinary(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinary", reflect.TypeOf((*MockUserPGRepository)(nil).AddBinary), ctx, userID, request)
}

// AddCard mocks base method.
func (m *MockUserPGRepository) AddCard(ctx context.Context, userID string, request *passbook.AddCardRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCard", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCard indicates an expected call of AddCard.
func (mr *MockUserPGRepositoryMockRecorder) AddCard(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockUserPGRepository)(nil).AddCard), ctx, userID, request)
}

// AddText mocks base method.
func (m *MockUserPGRepository) AddText(ctx context.Context, userID string, request *passbook.AddTextRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddText", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddText indicates an expected call of AddText.
func (mr *MockUserPGRepositoryMockRecorder) AddText(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddText", reflect.TypeOf((*MockUserPGRepository)(nil).AddText), ctx, userID, request)
}

// Create mocks base method.
func (m *MockUserPGRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserPGRepositoryMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserPGRepository)(nil).Create), ctx, user)
}

// FindByID mocks base method.
func (m *MockUserPGRepository) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, userID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserPGRepositoryMockRecorder) FindByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserPGRepository)(nil).FindByID), ctx, userID)
}

// FindByLogin mocks base method.
func (m *MockUserPGRepository) FindByLogin(ctx context.Context, login string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", ctx, login)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *MockUserPGRepositoryMockRecorder) FindByLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockUserPGRepository)(nil).FindByLogin), ctx, login)
}

// GetByTitle mocks base method.
func (m *MockUserPGRepository) GetByTitle(ctx context.Context, userID string, request *passbook.GetByTitleRequest) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTitle", ctx, userID, request)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTitle indicates an expected call of GetByTitle.
func (mr *MockUserPGRepositoryMockRecorder) GetByTitle(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTitle", reflect.TypeOf((*MockUserPGRepository)(nil).GetByTitle), ctx, userID, request)
}

// GetFullList mocks base method.
func (m *MockUserPGRepository) GetFullList(ctx context.Context, userID string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullList", ctx, userID)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullList indicates an expected call of GetFullList.
func (mr *MockUserPGRepositoryMockRecorder) GetFullList(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullList", reflect.TypeOf((*MockUserPGRepository)(nil).GetFullList), ctx, userID)
}
