// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/profile/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/profile/repository.go -destination=internal/infra/mocks/profile/mock_profile.go
//

// Package mock_profile is a generated GoMock package.
package mock_profile

import (
	context "context"
	reflect "reflect"

	profile "github.com/oprimogus/cardapiogo/internal/domain/profile"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateProfile mocks base method.
func (m *MockRepository) CreateProfile(ctx context.Context, userID string, params profile.CreateProfileParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", ctx, userID, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockRepositoryMockRecorder) CreateProfile(ctx, userID, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockRepository)(nil).CreateProfile), ctx, userID, params)
}