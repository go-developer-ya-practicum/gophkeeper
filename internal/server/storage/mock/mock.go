// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	models "github.com/go-developer-ya-practicum/gophkeeper/internal/server/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserStorage is a mock of UserStorage interface.
type MockUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUserStorageMockRecorder
}

// MockUserStorageMockRecorder is the mock recorder for MockUserStorage.
type MockUserStorageMockRecorder struct {
	mock *MockUserStorage
}

// NewMockUserStorage creates a new mock instance.
func NewMockUserStorage(ctrl *gomock.Controller) *MockUserStorage {
	mock := &MockUserStorage{ctrl: ctrl}
	mock.recorder = &MockUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorage) EXPECT() *MockUserStorageMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockUserStorage) GetUser(ctx context.Context, user *models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserStorageMockRecorder) GetUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserStorage)(nil).GetUser), ctx, user)
}

// PutUser mocks base method.
func (m *MockUserStorage) PutUser(ctx context.Context, user *models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutUser indicates an expected call of PutUser.
func (mr *MockUserStorageMockRecorder) PutUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUser", reflect.TypeOf((*MockUserStorage)(nil).PutUser), ctx, user)
}

// MockSecretStorage is a mock of SecretStorage interface.
type MockSecretStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSecretStorageMockRecorder
}

// MockSecretStorageMockRecorder is the mock recorder for MockSecretStorage.
type MockSecretStorageMockRecorder struct {
	mock *MockSecretStorage
}

// NewMockSecretStorage creates a new mock instance.
func NewMockSecretStorage(ctrl *gomock.Controller) *MockSecretStorage {
	mock := &MockSecretStorage{ctrl: ctrl}
	mock.recorder = &MockSecretStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretStorage) EXPECT() *MockSecretStorageMockRecorder {
	return m.recorder
}

// GetSecret mocks base method.
func (m *MockSecretStorage) GetSecret(ctx context.Context, name string, userID int) (*models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, name, userID)
	ret0, _ := ret[0].(*models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretStorageMockRecorder) GetSecret(ctx, name, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretStorage)(nil).GetSecret), ctx, name, userID)
}

// ListSecrets mocks base method.
func (m *MockSecretStorage) ListSecrets(ctx context.Context, userID int) ([]*models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", ctx, userID)
	ret0, _ := ret[0].([]*models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretStorageMockRecorder) ListSecrets(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretStorage)(nil).ListSecrets), ctx, userID)
}

// PutSecret mocks base method.
func (m *MockSecretStorage) PutSecret(ctx context.Context, secret *models.Secret) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSecret", ctx, secret)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutSecret indicates an expected call of PutSecret.
func (mr *MockSecretStorageMockRecorder) PutSecret(ctx, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSecret", reflect.TypeOf((*MockSecretStorage)(nil).PutSecret), ctx, secret)
}