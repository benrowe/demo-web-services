// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package mock_app is a generated GoMock package.
package mock_app

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEnvironment is a mock of Environment interface
type MockEnvironment struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentMockRecorder
}

// MockEnvironmentMockRecorder is the mock recorder for MockEnvironment
type MockEnvironmentMockRecorder struct {
	mock *MockEnvironment
}

// NewMockEnvironment creates a new mock instance
func NewMockEnvironment(ctrl *gomock.Controller) *MockEnvironment {
	mock := &MockEnvironment{ctrl: ctrl}
	mock.recorder = &MockEnvironmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironment) EXPECT() *MockEnvironmentMockRecorder {
	return m.recorder
}

// Getenv mocks base method
func (m *MockEnvironment) Getenv(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getenv", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Getenv indicates an expected call of Getenv
func (mr *MockEnvironmentMockRecorder) Getenv(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getenv", reflect.TypeOf((*MockEnvironment)(nil).Getenv), key)
}

// LookupEnv mocks base method
func (m *MockEnvironment) LookupEnv(key string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupEnv", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// LookupEnv indicates an expected call of LookupEnv
func (mr *MockEnvironmentMockRecorder) LookupEnv(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupEnv", reflect.TypeOf((*MockEnvironment)(nil).LookupEnv), key)
}