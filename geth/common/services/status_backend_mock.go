// Code generated by MockGen. DO NOT EDIT.
// Source: geth/common/services/status_backend.go

// Package services is a generated GoMock package.
package services

import (
	context "context"
	common "github.com/ethereum/go-ethereum/common"
	hexutil "github.com/ethereum/go-ethereum/common/hexutil"
	status "github.com/ethereum/go-ethereum/les/status"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStatusBackend is a mock of StatusBackend interface
type MockStatusBackend struct {
	ctrl     *gomock.Controller
	recorder *MockStatusBackendMockRecorder
}

// MockStatusBackendMockRecorder is the mock recorder for MockStatusBackend
type MockStatusBackendMockRecorder struct {
	mock *MockStatusBackend
}

// NewMockStatusBackend creates a new mock instance
func NewMockStatusBackend(ctrl *gomock.Controller) *MockStatusBackend {
	mock := &MockStatusBackend{ctrl: ctrl}
	mock.recorder = &MockStatusBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatusBackend) EXPECT() *MockStatusBackendMockRecorder {
	return m.recorder
}

// SetAccountsFilterHandler mocks base method
func (m *MockStatusBackend) SetAccountsFilterHandler(fn status.AccountsFilterHandler) {
	m.ctrl.Call(m, "SetAccountsFilterHandler", fn)
}

// SetAccountsFilterHandler indicates an expected call of SetAccountsFilterHandler
func (mr *MockStatusBackendMockRecorder) SetAccountsFilterHandler(fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccountsFilterHandler", reflect.TypeOf((*MockStatusBackend)(nil).SetAccountsFilterHandler), fn)
}

// AccountManager mocks base method
func (m *MockStatusBackend) AccountManager() *status.AccountManager {
	ret := m.ctrl.Call(m, "AccountManager")
	ret0, _ := ret[0].(*status.AccountManager)
	return ret0
}

// AccountManager indicates an expected call of AccountManager
func (mr *MockStatusBackendMockRecorder) AccountManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountManager", reflect.TypeOf((*MockStatusBackend)(nil).AccountManager))
}

// SendTransaction mocks base method
func (m *MockStatusBackend) SendTransaction(ctx context.Context, args status.SendTxArgs, passphrase string) (common.Hash, error) {
	ret := m.ctrl.Call(m, "SendTransaction", ctx, args, passphrase)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockStatusBackendMockRecorder) SendTransaction(ctx, args, passphrase interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockStatusBackend)(nil).SendTransaction), ctx, args, passphrase)
}

// EstimateGas mocks base method
func (m *MockStatusBackend) EstimateGas(ctx context.Context, args status.SendTxArgs) (*hexutil.Big, error) {
	ret := m.ctrl.Call(m, "EstimateGas", ctx, args)
	ret0, _ := ret[0].(*hexutil.Big)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas
func (mr *MockStatusBackendMockRecorder) EstimateGas(ctx, args interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockStatusBackend)(nil).EstimateGas), ctx, args)
}