// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUser is a mock of IUser interface.
type MockIUser struct {
	ctrl     *gomock.Controller
	recorder *MockIUserMockRecorder
}

// MockIUserMockRecorder is the mock recorder for MockIUser.
type MockIUserMockRecorder struct {
	mock *MockIUser
}

// NewMockIUser creates a new mock instance.
func NewMockIUser(ctrl *gomock.Controller) *MockIUser {
	mock := &MockIUser{ctrl: ctrl}
	mock.recorder = &MockIUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUser) EXPECT() *MockIUserMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIUser) Get(id string) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIUserMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUser)(nil).Get), id)
}