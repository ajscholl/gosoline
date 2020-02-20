// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import cfg "github.com/applike/gosoline/pkg/cfg"
import context "context"
import mock "github.com/stretchr/testify/mock"
import mon "github.com/applike/gosoline/pkg/mon"

// ConsumerCallback is an autogenerated mock type for the ConsumerCallback type
type ConsumerCallback struct {
	mock.Mock
}

// Boot provides a mock function with given fields: config, logger
func (_m *ConsumerCallback) Boot(config cfg.Config, logger mon.Logger) error {
	ret := _m.Called(config, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(cfg.Config, mon.Logger) error); ok {
		r0 = rf(config, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consume provides a mock function with given fields: ctx, model, attributes
func (_m *ConsumerCallback) Consume(ctx context.Context, model interface{}, attributes map[string]interface{}) (bool, error) {
	ret := _m.Called(ctx, model, attributes)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, map[string]interface{}) bool); ok {
		r0 = rf(ctx, model, attributes)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, map[string]interface{}) error); ok {
		r1 = rf(ctx, model, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModel provides a mock function with given fields:
func (_m *ConsumerCallback) GetModel() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}
