// Code generated by MockGen. DO NOT EDIT.
// Source: token_inspector.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	jwt "github.com/golang-jwt/jwt"
	gomock "github.com/golang/mock/gomock"
	jwt0 "github.com/quadev-ltd/qd-common/pkg/jwt"
	token "github.com/quadev-ltd/qd-common/pkg/token"
)

// MockTokenInspectorer is a mock of TokenInspectorer interface.
type MockTokenInspectorer struct {
	ctrl     *gomock.Controller
	recorder *MockTokenInspectorerMockRecorder
}

// MockTokenInspectorerMockRecorder is the mock recorder for MockTokenInspectorer.
type MockTokenInspectorerMockRecorder struct {
	mock *MockTokenInspectorer
}

// NewMockTokenInspectorer creates a new mock instance.
func NewMockTokenInspectorer(ctrl *gomock.Controller) *MockTokenInspectorer {
	mock := &MockTokenInspectorer{ctrl: ctrl}
	mock.recorder = &MockTokenInspectorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenInspectorer) EXPECT() *MockTokenInspectorerMockRecorder {
	return m.recorder
}

// GetClaimFromToken mocks base method.
func (m *MockTokenInspectorer) GetClaimFromToken(jwtToken *jwt.Token, claimKey string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaimFromToken", jwtToken, claimKey)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClaimFromToken indicates an expected call of GetClaimFromToken.
func (mr *MockTokenInspectorerMockRecorder) GetClaimFromToken(jwtToken, claimKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaimFromToken", reflect.TypeOf((*MockTokenInspectorer)(nil).GetClaimFromToken), jwtToken, claimKey)
}

// GetClaimsFromToken mocks base method.
func (m *MockTokenInspectorer) GetClaimsFromToken(token *jwt.Token) (*jwt0.TokenClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaimsFromToken", token)
	ret0, _ := ret[0].(*jwt0.TokenClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClaimsFromToken indicates an expected call of GetClaimsFromToken.
func (mr *MockTokenInspectorerMockRecorder) GetClaimsFromToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaimsFromToken", reflect.TypeOf((*MockTokenInspectorer)(nil).GetClaimsFromToken), token)
}

// GetEmailFromToken mocks base method.
func (m *MockTokenInspectorer) GetEmailFromToken(jwtToken *jwt.Token) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailFromToken", jwtToken)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailFromToken indicates an expected call of GetEmailFromToken.
func (mr *MockTokenInspectorerMockRecorder) GetEmailFromToken(jwtToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailFromToken", reflect.TypeOf((*MockTokenInspectorer)(nil).GetEmailFromToken), jwtToken)
}

// GetExpiryFromToken mocks base method.
func (m *MockTokenInspectorer) GetExpiryFromToken(jwtToken *jwt.Token) (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpiryFromToken", jwtToken)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpiryFromToken indicates an expected call of GetExpiryFromToken.
func (mr *MockTokenInspectorerMockRecorder) GetExpiryFromToken(jwtToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpiryFromToken", reflect.TypeOf((*MockTokenInspectorer)(nil).GetExpiryFromToken), jwtToken)
}

// GetTypeFromToken mocks base method.
func (m *MockTokenInspectorer) GetTypeFromToken(jwtToken *jwt.Token) (*token.TokenType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeFromToken", jwtToken)
	ret0, _ := ret[0].(*token.TokenType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypeFromToken indicates an expected call of GetTypeFromToken.
func (mr *MockTokenInspectorerMockRecorder) GetTypeFromToken(jwtToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeFromToken", reflect.TypeOf((*MockTokenInspectorer)(nil).GetTypeFromToken), jwtToken)
}

// GetUserIDFromToken mocks base method.
func (m *MockTokenInspectorer) GetUserIDFromToken(jwtToken *jwt.Token) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDFromToken", jwtToken)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDFromToken indicates an expected call of GetUserIDFromToken.
func (mr *MockTokenInspectorerMockRecorder) GetUserIDFromToken(jwtToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDFromToken", reflect.TypeOf((*MockTokenInspectorer)(nil).GetUserIDFromToken), jwtToken)
}
