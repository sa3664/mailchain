// Copyright 2021 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: receiver.go

// Package mailboxtest is a generated GoMock package.
package mailboxtest

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	stores "github.com/mailchain/mailchain/stores"
)

// MockReceiver is a mock of Receiver interface.
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverMockRecorder
}

// MockReceiverMockRecorder is the mock recorder for MockReceiver.
type MockReceiverMockRecorder struct {
	mock *MockReceiver
}

// NewMockReceiver creates a new mock instance.
func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiver) EXPECT() *MockReceiverMockRecorder {
	return m.recorder
}

// Kind mocks base method.
func (m *MockReceiver) Kind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockReceiverMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockReceiver)(nil).Kind))
}

// Receive mocks base method.
func (m *MockReceiver) Receive(ctx context.Context, protocol, network string, address []byte) ([]stores.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", ctx, protocol, network, address)
	ret0, _ := ret[0].([]stores.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockReceiverMockRecorder) Receive(ctx, protocol, network, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockReceiver)(nil).Receive), ctx, protocol, network, address)
}
