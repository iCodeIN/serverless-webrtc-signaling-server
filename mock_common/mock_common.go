// Code generated by MockGen. DO NOT EDIT.
// Source: common/interface.go

// Package mock_common is a generated GoMock package.
package mock_common

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApiGatewayManagementAPI is a mock of ApiGatewayManagementAPI interface
type MockApiGatewayManagementAPI struct {
	ctrl     *gomock.Controller
	recorder *MockApiGatewayManagementAPIMockRecorder
}

// MockApiGatewayManagementAPIMockRecorder is the mock recorder for MockApiGatewayManagementAPI
type MockApiGatewayManagementAPIMockRecorder struct {
	mock *MockApiGatewayManagementAPI
}

// NewMockApiGatewayManagementAPI creates a new mock instance
func NewMockApiGatewayManagementAPI(ctrl *gomock.Controller) *MockApiGatewayManagementAPI {
	mock := &MockApiGatewayManagementAPI{ctrl: ctrl}
	mock.recorder = &MockApiGatewayManagementAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiGatewayManagementAPI) EXPECT() *MockApiGatewayManagementAPIMockRecorder {
	return m.recorder
}

// PostToConnection mocks base method
func (m *MockApiGatewayManagementAPI) PostToConnection(connectionID, body string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostToConnection", connectionID, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostToConnection indicates an expected call of PostToConnection
func (mr *MockApiGatewayManagementAPIMockRecorder) PostToConnection(connectionID, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostToConnection", reflect.TypeOf((*MockApiGatewayManagementAPI)(nil).PostToConnection), connectionID, body)
}
