// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/bu_server/webhook/webhook_controller.go

// Package mock_webhook is a generated GoMock package.
package mock_webhook

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/openebl/openebl/pkg/bu_server/model"
	webhook "github.com/openebl/openebl/pkg/bu_server/webhook"
)

// MockWebhookController is a mock of WebhookController interface.
type MockWebhookController struct {
	ctrl     *gomock.Controller
	recorder *MockWebhookControllerMockRecorder
}

// MockWebhookControllerMockRecorder is the mock recorder for MockWebhookController.
type MockWebhookControllerMockRecorder struct {
	mock *MockWebhookController
}

// NewMockWebhookController creates a new mock instance.
func NewMockWebhookController(ctrl *gomock.Controller) *MockWebhookController {
	mock := &MockWebhookController{ctrl: ctrl}
	mock.recorder = &MockWebhookControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebhookController) EXPECT() *MockWebhookControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWebhookController) Create(ctx context.Context, ts int64, req webhook.CreateWebhookRequest) (model.Webhook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, ts, req)
	ret0, _ := ret[0].(model.Webhook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWebhookControllerMockRecorder) Create(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWebhookController)(nil).Create), ctx, ts, req)
}
