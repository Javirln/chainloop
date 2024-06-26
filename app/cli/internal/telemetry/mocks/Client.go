// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	telemetry "github.com/chainloop-dev/chainloop/app/cli/internal/telemetry"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// TrackEvent provides a mock function with given fields: ctx, eventName, id, tags
func (_m *Client) TrackEvent(ctx context.Context, eventName string, id string, tags telemetry.Tags) error {
	ret := _m.Called(ctx, eventName, id, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, telemetry.Tags) error); ok {
		r0 = rf(ctx, eventName, id, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
