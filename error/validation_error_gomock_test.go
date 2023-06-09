// Code generated by MockGen. DO NOT EDIT.
// Source: ./validation_error.go

// Package error is a generated GoMock package.
package error

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidationError is a mock of ValidationError interface.
type MockValidationError struct {
	ctrl     *gomock.Controller
	recorder *MockValidationErrorMockRecorder
}

// MockValidationErrorMockRecorder is the mock recorder for MockValidationError.
type MockValidationErrorMockRecorder struct {
	mock *MockValidationError
}

// NewMockValidationError creates a new mock instance.
func NewMockValidationError(ctrl *gomock.Controller) *MockValidationError {
	mock := &MockValidationError{ctrl: ctrl}
	mock.recorder = &MockValidationErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidationError) EXPECT() *MockValidationErrorMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockValidationError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockValidationErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockValidationError)(nil).Error))
}

// GetRule mocks base method.
func (m *MockValidationError) GetRule() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRule")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRule indicates an expected call of GetRule.
func (mr *MockValidationErrorMockRecorder) GetRule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockValidationError)(nil).GetRule))
}
