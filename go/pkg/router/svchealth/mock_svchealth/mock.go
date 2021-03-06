// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/pkg/router/svchealth (interfaces: Discoverer)

// Package mock_svchealth is a generated GoMock package.
package mock_svchealth

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	addr "github.com/scionproto/scion/go/lib/addr"
	net "net"
	reflect "reflect"
)

// MockDiscoverer is a mock of Discoverer interface
type MockDiscoverer struct {
	ctrl     *gomock.Controller
	recorder *MockDiscovererMockRecorder
}

// MockDiscovererMockRecorder is the mock recorder for MockDiscoverer
type MockDiscovererMockRecorder struct {
	mock *MockDiscoverer
}

// NewMockDiscoverer creates a new mock instance
func NewMockDiscoverer(ctrl *gomock.Controller) *MockDiscoverer {
	mock := &MockDiscoverer{ctrl: ctrl}
	mock.recorder = &MockDiscovererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoverer) EXPECT() *MockDiscovererMockRecorder {
	return m.recorder
}

// Discover mocks base method
func (m *MockDiscoverer) Discover(arg0 context.Context, arg1 addr.HostSVC) ([]*net.UDPAddr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discover", arg0, arg1)
	ret0, _ := ret[0].([]*net.UDPAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Discover indicates an expected call of Discover
func (mr *MockDiscovererMockRecorder) Discover(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discover", reflect.TypeOf((*MockDiscoverer)(nil).Discover), arg0, arg1)
}

// Discoverable mocks base method
func (m *MockDiscoverer) Discoverable(arg0 addr.HostSVC) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discoverable", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Discoverable indicates an expected call of Discoverable
func (mr *MockDiscovererMockRecorder) Discoverable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discoverable", reflect.TypeOf((*MockDiscoverer)(nil).Discoverable), arg0)
}
