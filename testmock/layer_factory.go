// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/lifecycle (interfaces: LayerFactory)

// Package testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	launch "github.com/buildpacks/lifecycle/launch"
	layers "github.com/buildpacks/lifecycle/layers"
)

// MockLayerFactory is a mock of LayerFactory interface
type MockLayerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockLayerFactoryMockRecorder
}

// MockLayerFactoryMockRecorder is the mock recorder for MockLayerFactory
type MockLayerFactoryMockRecorder struct {
	mock *MockLayerFactory
}

// NewMockLayerFactory creates a new mock instance
func NewMockLayerFactory(ctrl *gomock.Controller) *MockLayerFactory {
	mock := &MockLayerFactory{ctrl: ctrl}
	mock.recorder = &MockLayerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLayerFactory) EXPECT() *MockLayerFactoryMockRecorder {
	return m.recorder
}

// DirLayer mocks base method
func (m *MockLayerFactory) DirLayer(arg0, arg1 string) (layers.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DirLayer", arg0, arg1)
	ret0, _ := ret[0].(layers.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DirLayer indicates an expected call of DirLayer
func (mr *MockLayerFactoryMockRecorder) DirLayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DirLayer", reflect.TypeOf((*MockLayerFactory)(nil).DirLayer), arg0, arg1)
}

// LauncherLayer mocks base method
func (m *MockLayerFactory) LauncherLayer(arg0 string) (layers.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LauncherLayer", arg0)
	ret0, _ := ret[0].(layers.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LauncherLayer indicates an expected call of LauncherLayer
func (mr *MockLayerFactoryMockRecorder) LauncherLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LauncherLayer", reflect.TypeOf((*MockLayerFactory)(nil).LauncherLayer), arg0)
}

// ProcessTypesLayer mocks base method
func (m *MockLayerFactory) ProcessTypesLayer(arg0 launch.Metadata) (layers.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTypesLayer", arg0)
	ret0, _ := ret[0].(layers.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTypesLayer indicates an expected call of ProcessTypesLayer
func (mr *MockLayerFactoryMockRecorder) ProcessTypesLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTypesLayer", reflect.TypeOf((*MockLayerFactory)(nil).ProcessTypesLayer), arg0)
}

// SliceLayers mocks base method
func (m *MockLayerFactory) SliceLayers(arg0 string, arg1 []layers.Slice) ([]layers.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SliceLayers", arg0, arg1)
	ret0, _ := ret[0].([]layers.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SliceLayers indicates an expected call of SliceLayers
func (mr *MockLayerFactoryMockRecorder) SliceLayers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SliceLayers", reflect.TypeOf((*MockLayerFactory)(nil).SliceLayers), arg0, arg1)
}
