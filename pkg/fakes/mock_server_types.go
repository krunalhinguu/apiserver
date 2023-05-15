// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/apiserver/pkg/types (interfaces: ResponseWriter,AccessControl)

// Package fakes is a generated GoMock package.
package fakes

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/rancher/apiserver/pkg/types"
)

// MockResponseWriter is a mock of ResponseWriter interface.
type MockResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockResponseWriterMockRecorder
}

// MockResponseWriterMockRecorder is the mock recorder for MockResponseWriter.
type MockResponseWriterMockRecorder struct {
	mock *MockResponseWriter
}

// NewMockResponseWriter creates a new mock instance.
func NewMockResponseWriter(ctrl *gomock.Controller) *MockResponseWriter {
	mock := &MockResponseWriter{ctrl: ctrl}
	mock.recorder = &MockResponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponseWriter) EXPECT() *MockResponseWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockResponseWriter) Write(arg0 *types.APIRequest, arg1 int, arg2 types.APIObject) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", arg0, arg1, arg2)
}

// Write indicates an expected call of Write.
func (mr *MockResponseWriterMockRecorder) Write(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockResponseWriter)(nil).Write), arg0, arg1, arg2)
}

// WriteList mocks base method.
func (m *MockResponseWriter) WriteList(arg0 *types.APIRequest, arg1 int, arg2 types.APIObjectList) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteList", arg0, arg1, arg2)
}

// WriteList indicates an expected call of WriteList.
func (mr *MockResponseWriterMockRecorder) WriteList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteList", reflect.TypeOf((*MockResponseWriter)(nil).WriteList), arg0, arg1, arg2)
}

// MockAccessControl is a mock of AccessControl interface.
type MockAccessControl struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlMockRecorder
}

// MockAccessControlMockRecorder is the mock recorder for MockAccessControl.
type MockAccessControlMockRecorder struct {
	mock *MockAccessControl
}

// NewMockAccessControl creates a new mock instance.
func NewMockAccessControl(ctrl *gomock.Controller) *MockAccessControl {
	mock := &MockAccessControl{ctrl: ctrl}
	mock.recorder = &MockAccessControlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControl) EXPECT() *MockAccessControlMockRecorder {
	return m.recorder
}

// CanAction mocks base method.
func (m *MockAccessControl) CanAction(arg0 *types.APIRequest, arg1 *types.APISchema, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanAction", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanAction indicates an expected call of CanAction.
func (mr *MockAccessControlMockRecorder) CanAction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanAction", reflect.TypeOf((*MockAccessControl)(nil).CanAction), arg0, arg1, arg2)
}

// CanCreate mocks base method.
func (m *MockAccessControl) CanCreate(arg0 *types.APIRequest, arg1 *types.APISchema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanCreate indicates an expected call of CanCreate.
func (mr *MockAccessControlMockRecorder) CanCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreate", reflect.TypeOf((*MockAccessControl)(nil).CanCreate), arg0, arg1)
}

// CanDelete mocks base method.
func (m *MockAccessControl) CanDelete(arg0 *types.APIRequest, arg1 types.APIObject, arg2 *types.APISchema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDelete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanDelete indicates an expected call of CanDelete.
func (mr *MockAccessControlMockRecorder) CanDelete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDelete", reflect.TypeOf((*MockAccessControl)(nil).CanDelete), arg0, arg1, arg2)
}

// CanDo mocks base method.
func (m *MockAccessControl) CanDo(arg0 *types.APIRequest, arg1, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDo", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanDo indicates an expected call of CanDo.
func (mr *MockAccessControlMockRecorder) CanDo(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDo", reflect.TypeOf((*MockAccessControl)(nil).CanDo), arg0, arg1, arg2, arg3, arg4)
}

// CanGet mocks base method.
func (m *MockAccessControl) CanGet(arg0 *types.APIRequest, arg1 *types.APISchema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanGet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanGet indicates an expected call of CanGet.
func (mr *MockAccessControlMockRecorder) CanGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanGet", reflect.TypeOf((*MockAccessControl)(nil).CanGet), arg0, arg1)
}

// CanList mocks base method.
func (m *MockAccessControl) CanList(arg0 *types.APIRequest, arg1 *types.APISchema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanList indicates an expected call of CanList.
func (mr *MockAccessControlMockRecorder) CanList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanList", reflect.TypeOf((*MockAccessControl)(nil).CanList), arg0, arg1)
}

// CanUpdate mocks base method.
func (m *MockAccessControl) CanUpdate(arg0 *types.APIRequest, arg1 types.APIObject, arg2 *types.APISchema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanUpdate indicates an expected call of CanUpdate.
func (mr *MockAccessControlMockRecorder) CanUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanUpdate", reflect.TypeOf((*MockAccessControl)(nil).CanUpdate), arg0, arg1, arg2)
}

// CanWatch mocks base method.
func (m *MockAccessControl) CanWatch(arg0 *types.APIRequest, arg1 *types.APISchema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanWatch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanWatch indicates an expected call of CanWatch.
func (mr *MockAccessControlMockRecorder) CanWatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanWatch", reflect.TypeOf((*MockAccessControl)(nil).CanWatch), arg0, arg1)
}