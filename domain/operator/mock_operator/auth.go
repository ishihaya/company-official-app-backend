// Code generated by MockGen. DO NOT EDIT.
// Source: domain/operator/auth.go

// Package mock_operator is a generated GoMock package.
package mock_operator

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/ishihaya/company-official-app-backend/domain/entity"
)

// MockAuthOperator is a mock of AuthOperator interface.
type MockAuthOperator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthOperatorMockRecorder
}

// MockAuthOperatorMockRecorder is the mock recorder for MockAuthOperator.
type MockAuthOperatorMockRecorder struct {
	mock *MockAuthOperator
}

// NewMockAuthOperator creates a new mock instance.
func NewMockAuthOperator(ctrl *gomock.Controller) *MockAuthOperator {
	mock := &MockAuthOperator{ctrl: ctrl}
	mock.recorder = &MockAuthOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthOperator) EXPECT() *MockAuthOperatorMockRecorder {
	return m.recorder
}

// FindByToken mocks base method.
func (m *MockAuthOperator) FindByToken(token string) (*entity.Auth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", token)
	ret0, _ := ret[0].(*entity.Auth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken.
func (mr *MockAuthOperatorMockRecorder) FindByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockAuthOperator)(nil).FindByToken), token)
}
