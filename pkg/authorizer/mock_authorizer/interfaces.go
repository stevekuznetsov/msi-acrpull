// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/authorizer/interfaces.go
//
// Generated by this command:
//
//	mockgen-v0.4.0 -source=pkg/authorizer/interfaces.go
//

// Package mock_authorizer is a generated GoMock package.
package mock_authorizer

import (
	context "context"
	reflect "reflect"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AcquireACRAccessTokenWithClientID mocks base method.
func (m *MockInterface) AcquireACRAccessTokenWithClientID(ctx context.Context, clientID, acrFQDN string) (azcore.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireACRAccessTokenWithClientID", ctx, clientID, acrFQDN)
	ret0, _ := ret[0].(azcore.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireACRAccessTokenWithClientID indicates an expected call of AcquireACRAccessTokenWithClientID.
func (mr *MockInterfaceMockRecorder) AcquireACRAccessTokenWithClientID(ctx, clientID, acrFQDN any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireACRAccessTokenWithClientID", reflect.TypeOf((*MockInterface)(nil).AcquireACRAccessTokenWithClientID), ctx, clientID, acrFQDN)
}

// AcquireACRAccessTokenWithResourceID mocks base method.
func (m *MockInterface) AcquireACRAccessTokenWithResourceID(ctx context.Context, identityResourceID, acrFQDN string) (azcore.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireACRAccessTokenWithResourceID", ctx, identityResourceID, acrFQDN)
	ret0, _ := ret[0].(azcore.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireACRAccessTokenWithResourceID indicates an expected call of AcquireACRAccessTokenWithResourceID.
func (mr *MockInterfaceMockRecorder) AcquireACRAccessTokenWithResourceID(ctx, identityResourceID, acrFQDN any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireACRAccessTokenWithResourceID", reflect.TypeOf((*MockInterface)(nil).AcquireACRAccessTokenWithResourceID), ctx, identityResourceID, acrFQDN)
}
