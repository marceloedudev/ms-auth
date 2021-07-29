// Code generated by MockGen. DO NOT EDIT.
// Source: internal/client/usecase/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	domain "ms-auth/internal/client/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPgRepository is a mock of PgRepository interface.
type MockPgRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPgRepositoryMockRecorder
}

// MockPgRepositoryMockRecorder is the mock recorder for MockPgRepository.
type MockPgRepositoryMockRecorder struct {
	mock *MockPgRepository
}

// NewMockPgRepository creates a new mock instance.
func NewMockPgRepository(ctrl *gomock.Controller) *MockPgRepository {
	mock := &MockPgRepository{ctrl: ctrl}
	mock.recorder = &MockPgRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPgRepository) EXPECT() *MockPgRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPgRepository) Create(ctx context.Context, e *domain.Client) (*domain.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, e)
	ret0, _ := ret[0].(*domain.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPgRepositoryMockRecorder) Create(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPgRepository)(nil).Create), ctx, e)
}

// FindByText mocks base method.
func (m *MockPgRepository) FindByText(ctx context.Context, text string, userID int64) (*domain.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByText", ctx, text, userID)
	ret0, _ := ret[0].(*domain.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByText indicates an expected call of FindByText.
func (mr *MockPgRepositoryMockRecorder) FindByText(ctx, text, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByText", reflect.TypeOf((*MockPgRepository)(nil).FindByText), ctx, text, userID)
}

// FindClientByIDAndSecret mocks base method.
func (m *MockPgRepository) FindClientByIDAndSecret(ctx context.Context, clientID, clientSecret string) (*domain.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindClientByIDAndSecret", ctx, clientID, clientSecret)
	ret0, _ := ret[0].(*domain.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindClientByIDAndSecret indicates an expected call of FindClientByIDAndSecret.
func (mr *MockPgRepositoryMockRecorder) FindClientByIDAndSecret(ctx, clientID, clientSecret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindClientByIDAndSecret", reflect.TypeOf((*MockPgRepository)(nil).FindClientByIDAndSecret), ctx, clientID, clientSecret)
}

// FindManagingByText mocks base method.
func (m *MockPgRepository) FindManagingByText(ctx context.Context, text string) (*domain.Managing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindManagingByText", ctx, text)
	ret0, _ := ret[0].(*domain.Managing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindManagingByText indicates an expected call of FindManagingByText.
func (mr *MockPgRepositoryMockRecorder) FindManagingByText(ctx, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindManagingByText", reflect.TypeOf((*MockPgRepository)(nil).FindManagingByText), ctx, text)
}

// MockUseCaseClient is a mock of UseCaseClient interface.
type MockUseCaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseClientMockRecorder
}

// MockUseCaseClientMockRecorder is the mock recorder for MockUseCaseClient.
type MockUseCaseClientMockRecorder struct {
	mock *MockUseCaseClient
}

// NewMockUseCaseClient creates a new mock instance.
func NewMockUseCaseClient(ctrl *gomock.Controller) *MockUseCaseClient {
	mock := &MockUseCaseClient{ctrl: ctrl}
	mock.recorder = &MockUseCaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCaseClient) EXPECT() *MockUseCaseClientMockRecorder {
	return m.recorder
}

// RegisterClient mocks base method.
func (m *MockUseCaseClient) RegisterClient(ctx context.Context, clientID, clientSecret, managingID string, userID int64) (*domain.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterClient", ctx, clientID, clientSecret, managingID, userID)
	ret0, _ := ret[0].(*domain.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterClient indicates an expected call of RegisterClient.
func (mr *MockUseCaseClientMockRecorder) RegisterClient(ctx, clientID, clientSecret, managingID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterClient", reflect.TypeOf((*MockUseCaseClient)(nil).RegisterClient), ctx, clientID, clientSecret, managingID, userID)
}

// VerifyClient mocks base method.
func (m *MockUseCaseClient) VerifyClient(ctx context.Context, clientID, clientSecret string) (*domain.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyClient", ctx, clientID, clientSecret)
	ret0, _ := ret[0].(*domain.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyClient indicates an expected call of VerifyClient.
func (mr *MockUseCaseClientMockRecorder) VerifyClient(ctx, clientID, clientSecret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyClient", reflect.TypeOf((*MockUseCaseClient)(nil).VerifyClient), ctx, clientID, clientSecret)
}

// VerifyManaging mocks base method.
func (m *MockUseCaseClient) VerifyManaging(ctx context.Context, text string) (*domain.Managing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyManaging", ctx, text)
	ret0, _ := ret[0].(*domain.Managing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyManaging indicates an expected call of VerifyManaging.
func (mr *MockUseCaseClientMockRecorder) VerifyManaging(ctx, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyManaging", reflect.TypeOf((*MockUseCaseClient)(nil).VerifyManaging), ctx, text)
}