// Code generated by MockGen. DO NOT EDIT.
// Source: subscriber/subscriber.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "pub-sub-service/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSubscriber is a mock of Subscriber interface.
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberMockRecorder
}

// MockSubscriberMockRecorder is the mock recorder for MockSubscriber.
type MockSubscriberMockRecorder struct {
	mock *MockSubscriber
}

// NewMockSubscriber creates a new mock instance.
func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &MockSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriber) EXPECT() *MockSubscriberMockRecorder {
	return m.recorder
}

// AddEvent mocks base method.
func (m *MockSubscriber) AddEvent(event model.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEvent", event)
}

// AddEvent indicates an expected call of AddEvent.
func (mr *MockSubscriberMockRecorder) AddEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEvent", reflect.TypeOf((*MockSubscriber)(nil).AddEvent), event)
}

// Process mocks base method.
func (m *MockSubscriber) Process() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Process")
}

// Process indicates an expected call of Process.
func (mr *MockSubscriberMockRecorder) Process() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockSubscriber)(nil).Process))
}
