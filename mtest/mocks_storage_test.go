// Code generated by MockGen. DO NOT EDIT.
// Source: ./storage/storage.go
//
// Generated by this command:
//
//	mockgen -source ./storage/storage.go -package mtest
//

// Package mtest is a generated GoMock package.
package mtest

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
	isgomock struct{}
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProduct) Create(description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", description)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProductMockRecorder) Create(description any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProduct)(nil).Create), description)
}

// Description mocks base method.
func (m *MockProduct) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description.
func (mr *MockProductMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockProduct)(nil).Description))
}
