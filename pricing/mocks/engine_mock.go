// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vegaprotocol/priceproxy/pricing (interfaces: Engine)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	config "github.com/vegaprotocol/priceproxy/config"
	pricing "github.com/vegaprotocol/priceproxy/pricing"
	reflect "reflect"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// AddPrice mocks base method
func (m *MockEngine) AddPrice(arg0 config.PriceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPrice", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPrice indicates an expected call of AddPrice
func (mr *MockEngineMockRecorder) AddPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPrice", reflect.TypeOf((*MockEngine)(nil).AddPrice), arg0)
}

// AddSource mocks base method
func (m *MockEngine) AddSource(arg0 config.SourceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSource", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSource indicates an expected call of AddSource
func (mr *MockEngineMockRecorder) AddSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSource", reflect.TypeOf((*MockEngine)(nil).AddSource), arg0)
}

// GetPrice mocks base method
func (m *MockEngine) GetPrice(arg0 config.PriceConfig) (pricing.PriceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice", arg0)
	ret0, _ := ret[0].(pricing.PriceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrice indicates an expected call of GetPrice
func (mr *MockEngineMockRecorder) GetPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockEngine)(nil).GetPrice), arg0)
}

// GetPrices mocks base method
func (m *MockEngine) GetPrices() map[config.PriceConfig]pricing.PriceInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrices")
	ret0, _ := ret[0].(map[config.PriceConfig]pricing.PriceInfo)
	return ret0
}

// GetPrices indicates an expected call of GetPrices
func (mr *MockEngineMockRecorder) GetPrices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrices", reflect.TypeOf((*MockEngine)(nil).GetPrices))
}

// GetSource mocks base method
func (m *MockEngine) GetSource(arg0 string) (config.SourceConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSource", arg0)
	ret0, _ := ret[0].(config.SourceConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSource indicates an expected call of GetSource
func (mr *MockEngineMockRecorder) GetSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSource", reflect.TypeOf((*MockEngine)(nil).GetSource), arg0)
}

// GetSources mocks base method
func (m *MockEngine) GetSources() ([]config.SourceConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSources")
	ret0, _ := ret[0].([]config.SourceConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSources indicates an expected call of GetSources
func (mr *MockEngineMockRecorder) GetSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSources", reflect.TypeOf((*MockEngine)(nil).GetSources))
}

// UpdatePrice mocks base method
func (m *MockEngine) UpdatePrice(arg0 config.PriceConfig, arg1 pricing.PriceInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePrice", arg0, arg1)
}

// UpdatePrice indicates an expected call of UpdatePrice
func (mr *MockEngineMockRecorder) UpdatePrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePrice", reflect.TypeOf((*MockEngine)(nil).UpdatePrice), arg0, arg1)
}

// WaitForPrice mocks base method
func (m *MockEngine) WaitForPrice(arg0 config.PriceConfig) pricing.PriceInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForPrice", arg0)
	ret0, _ := ret[0].(pricing.PriceInfo)
	return ret0
}

// WaitForPrice indicates an expected call of WaitForPrice
func (mr *MockEngineMockRecorder) WaitForPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForPrice", reflect.TypeOf((*MockEngine)(nil).WaitForPrice), arg0)
}
