// Automatically generated by MockGen. DO NOT EDIT!
// Source: storage.go

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
)

// Mock of CodeResponseTypeStorage interface
type MockCodeResponseTypeStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockCodeResponseTypeStorageRecorder
}

// Recorder for MockCodeResponseTypeStorage (not exported)
type _MockCodeResponseTypeStorageRecorder struct {
	mock *MockCodeResponseTypeStorage
}

func NewMockCodeResponseTypeStorage(ctrl *gomock.Controller) *MockCodeResponseTypeStorage {
	mock := &MockCodeResponseTypeStorage{ctrl: ctrl}
	mock.recorder = &_MockCodeResponseTypeStorageRecorder{mock}
	return mock
}

func (_m *MockCodeResponseTypeStorage) EXPECT() *_MockCodeResponseTypeStorageRecorder {
	return _m.recorder
}

func (_m *MockCodeResponseTypeStorage) CreateAuthorizeCodeSession(code string, authorizeRequest fosite.AuthorizeRequester, extra interface{}) error {
	ret := _m.ctrl.Call(_m, "CreateAuthorizeCodeSession", code, authorizeRequest, extra)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCodeResponseTypeStorageRecorder) CreateAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockCodeResponseTypeStorage) GetAuthorizeCodeSession(code string, authorizeRequest fosite.AuthorizeRequester, extra interface{}) error {
	ret := _m.ctrl.Call(_m, "GetAuthorizeCodeSession", code, authorizeRequest, extra)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCodeResponseTypeStorageRecorder) GetAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockCodeResponseTypeStorage) DeleteAuthorizeCodeSession(code string) error {
	ret := _m.ctrl.Call(_m, "DeleteAuthorizeCodeSession", code)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCodeResponseTypeStorageRecorder) DeleteAuthorizeCodeSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAuthorizeCodeSession", arg0)
}
