// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/app/app.go
//
// Generated by this command:
//
//	mockgen -package=app -source=pkg/app/app.go
//

// Package app is a generated GoMock package.
package app

import (
	reflect "reflect"

	context "github.com/free5gc/smf/internal/context"
	factory "github.com/free5gc/smf/pkg/factory"
	gomock "go.uber.org/mock/gomock"
)

// MockApp is a mock of App interface.
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *MockAppMockRecorder
}

// MockAppMockRecorder is the mock recorder for MockApp.
type MockAppMockRecorder struct {
	mock *MockApp
}

// NewMockApp creates a new mock instance.
func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &MockAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApp) EXPECT() *MockAppMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockApp) Config() *factory.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*factory.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockAppMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockApp)(nil).Config))
}

// Context mocks base method.
func (m *MockApp) Context() *context.SMFContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(*context.SMFContext)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAppMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockApp)(nil).Context))
}

// SetLogEnable mocks base method.
func (m *MockApp) SetLogEnable(enable bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogEnable", enable)
}

// SetLogEnable indicates an expected call of SetLogEnable.
func (mr *MockAppMockRecorder) SetLogEnable(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogEnable", reflect.TypeOf((*MockApp)(nil).SetLogEnable), enable)
}

// SetLogLevel mocks base method.
func (m *MockApp) SetLogLevel(level string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogLevel", level)
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockAppMockRecorder) SetLogLevel(level any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockApp)(nil).SetLogLevel), level)
}

// SetReportCaller mocks base method.
func (m *MockApp) SetReportCaller(reportCaller bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReportCaller", reportCaller)
}

// SetReportCaller indicates an expected call of SetReportCaller.
func (mr *MockAppMockRecorder) SetReportCaller(reportCaller any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReportCaller", reflect.TypeOf((*MockApp)(nil).SetReportCaller), reportCaller)
}

// Start mocks base method.
func (m *MockApp) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockAppMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockApp)(nil).Start))
}

// Terminate mocks base method.
func (m *MockApp) Terminate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Terminate")
}

// Terminate indicates an expected call of Terminate.
func (mr *MockAppMockRecorder) Terminate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terminate", reflect.TypeOf((*MockApp)(nil).Terminate))
}