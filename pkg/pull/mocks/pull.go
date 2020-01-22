// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/wasme/pkg/pull (interfaces: ImagePuller)

// Package mock_pull is a generated GoMock package.
package mock_pull

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	config "github.com/solo-io/wasme/pkg/config"
	pull "github.com/solo-io/wasme/pkg/pull"
)

// MockImagePuller is a mock of ImagePuller interface
type MockImagePuller struct {
	ctrl     *gomock.Controller
	recorder *MockImagePullerMockRecorder
}

// MockImagePullerMockRecorder is the mock recorder for MockImagePuller
type MockImagePullerMockRecorder struct {
	mock *MockImagePuller
}

// NewMockImagePuller creates a new mock instance
func NewMockImagePuller(ctrl *gomock.Controller) *MockImagePuller {
	mock := &MockImagePuller{ctrl: ctrl}
	mock.recorder = &MockImagePullerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImagePuller) EXPECT() *MockImagePullerMockRecorder {
	return m.recorder
}

// Pull mocks base method
func (m *MockImagePuller) Pull(arg0 context.Context, arg1 string) ([]v1.Descriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", arg0, arg1)
	ret0, _ := ret[0].([]v1.Descriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pull indicates an expected call of Pull
func (mr *MockImagePullerMockRecorder) Pull(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockImagePuller)(nil).Pull), arg0, arg1)
}

// PullCodeDescriptor mocks base method
func (m *MockImagePuller) PullCodeDescriptor(arg0 context.Context, arg1 string) (v1.Descriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullCodeDescriptor", arg0, arg1)
	ret0, _ := ret[0].(v1.Descriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullCodeDescriptor indicates an expected call of PullCodeDescriptor
func (mr *MockImagePullerMockRecorder) PullCodeDescriptor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullCodeDescriptor", reflect.TypeOf((*MockImagePuller)(nil).PullCodeDescriptor), arg0, arg1)
}

// PullConfigFile mocks base method
func (m *MockImagePuller) PullConfigFile(arg0 context.Context, arg1 string) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullConfigFile", arg0, arg1)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullConfigFile indicates an expected call of PullConfigFile
func (mr *MockImagePullerMockRecorder) PullConfigFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullConfigFile", reflect.TypeOf((*MockImagePuller)(nil).PullConfigFile), arg0, arg1)
}

// PullFilter mocks base method
func (m *MockImagePuller) PullFilter(arg0 context.Context, arg1 string) (pull.Filter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullFilter", arg0, arg1)
	ret0, _ := ret[0].(pull.Filter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullFilter indicates an expected call of PullFilter
func (mr *MockImagePullerMockRecorder) PullFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullFilter", reflect.TypeOf((*MockImagePuller)(nil).PullFilter), arg0, arg1)
}
