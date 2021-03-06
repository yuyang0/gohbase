// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yuyang0/gohbase (interfaces: AdminClient)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	hrpc "github.com/yuyang0/gohbase/hrpc"
	pb "github.com/yuyang0/gohbase/pb"
	reflect "reflect"
)

// MockAdminClient is a mock of AdminClient interface
type MockAdminClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminClientMockRecorder
}

// MockAdminClientMockRecorder is the mock recorder for MockAdminClient
type MockAdminClientMockRecorder struct {
	mock *MockAdminClient
}

// NewMockAdminClient creates a new mock instance
func NewMockAdminClient(ctrl *gomock.Controller) *MockAdminClient {
	mock := &MockAdminClient{ctrl: ctrl}
	mock.recorder = &MockAdminClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminClient) EXPECT() *MockAdminClientMockRecorder {
	return m.recorder
}

// ClusterStatus mocks base method
func (m *MockAdminClient) ClusterStatus() (*pb.ClusterStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterStatus")
	ret0, _ := ret[0].(*pb.ClusterStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterStatus indicates an expected call of ClusterStatus
func (mr *MockAdminClientMockRecorder) ClusterStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterStatus", reflect.TypeOf((*MockAdminClient)(nil).ClusterStatus))
}

// CreateNamespace mocks base method
func (m *MockAdminClient) CreateNamespace(arg0 *hrpc.CreateNamespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNamespace indicates an expected call of CreateNamespace
func (mr *MockAdminClientMockRecorder) CreateNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockAdminClient)(nil).CreateNamespace), arg0)
}

// CreateTable mocks base method
func (m *MockAdminClient) CreateTable(arg0 *hrpc.CreateTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTable indicates an expected call of CreateTable
func (mr *MockAdminClientMockRecorder) CreateTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockAdminClient)(nil).CreateTable), arg0)
}

// DeleteNamespace mocks base method
func (m *MockAdminClient) DeleteNamespace(arg0 *hrpc.DeleteNamespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespace indicates an expected call of DeleteNamespace
func (mr *MockAdminClientMockRecorder) DeleteNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockAdminClient)(nil).DeleteNamespace), arg0)
}

// DeleteTable mocks base method
func (m *MockAdminClient) DeleteTable(arg0 *hrpc.DeleteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTable indicates an expected call of DeleteTable
func (mr *MockAdminClientMockRecorder) DeleteTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTable", reflect.TypeOf((*MockAdminClient)(nil).DeleteTable), arg0)
}

// DisableTable mocks base method
func (m *MockAdminClient) DisableTable(arg0 *hrpc.DisableTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableTable indicates an expected call of DisableTable
func (mr *MockAdminClientMockRecorder) DisableTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableTable", reflect.TypeOf((*MockAdminClient)(nil).DisableTable), arg0)
}

// EnableTable mocks base method
func (m *MockAdminClient) EnableTable(arg0 *hrpc.EnableTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableTable indicates an expected call of EnableTable
func (mr *MockAdminClientMockRecorder) EnableTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableTable", reflect.TypeOf((*MockAdminClient)(nil).EnableTable), arg0)
}

// ListNamespace mocks base method
func (m *MockAdminClient) ListNamespace(arg0 *hrpc.ListNamespace) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespace", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespace indicates an expected call of ListNamespace
func (mr *MockAdminClientMockRecorder) ListNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespace", reflect.TypeOf((*MockAdminClient)(nil).ListNamespace), arg0)
}

// ListTableNames mocks base method
func (m *MockAdminClient) ListTableNames(arg0 *hrpc.ListTableName) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTableNames indicates an expected call of ListTableNames
func (mr *MockAdminClientMockRecorder) ListTableNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableNames", reflect.TypeOf((*MockAdminClient)(nil).ListTableNames), arg0)
}
