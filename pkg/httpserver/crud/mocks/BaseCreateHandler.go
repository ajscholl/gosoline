// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mock "github.com/stretchr/testify/mock"
)

// BaseCreateHandler is an autogenerated mock type for the BaseCreateHandler type
type BaseCreateHandler struct {
	mock.Mock
}

type BaseCreateHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *BaseCreateHandler) EXPECT() *BaseCreateHandler_Expecter {
	return &BaseCreateHandler_Expecter{mock: &_m.Mock}
}

// GetCreateInput provides a mock function with given fields:
func (_m *BaseCreateHandler) GetCreateInput() interface{} {
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

// BaseCreateHandler_GetCreateInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCreateInput'
type BaseCreateHandler_GetCreateInput_Call struct {
	*mock.Call
}

// GetCreateInput is a helper method to define mock.On call
func (_e *BaseCreateHandler_Expecter) GetCreateInput() *BaseCreateHandler_GetCreateInput_Call {
	return &BaseCreateHandler_GetCreateInput_Call{Call: _e.mock.On("GetCreateInput")}
}

func (_c *BaseCreateHandler_GetCreateInput_Call) Run(run func()) *BaseCreateHandler_GetCreateInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BaseCreateHandler_GetCreateInput_Call) Return(_a0 interface{}) *BaseCreateHandler_GetCreateInput_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BaseCreateHandler_GetCreateInput_Call) RunAndReturn(run func() interface{}) *BaseCreateHandler_GetCreateInput_Call {
	_c.Call.Return(run)
	return _c
}

// TransformCreate provides a mock function with given fields: ctx, input, model
func (_m *BaseCreateHandler) TransformCreate(ctx context.Context, input interface{}, model db_repo.ModelBased) error {
	ret := _m.Called(ctx, input, model)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, input, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BaseCreateHandler_TransformCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransformCreate'
type BaseCreateHandler_TransformCreate_Call struct {
	*mock.Call
}

// TransformCreate is a helper method to define mock.On call
//   - ctx context.Context
//   - input interface{}
//   - model db_repo.ModelBased
func (_e *BaseCreateHandler_Expecter) TransformCreate(ctx interface{}, input interface{}, model interface{}) *BaseCreateHandler_TransformCreate_Call {
	return &BaseCreateHandler_TransformCreate_Call{Call: _e.mock.On("TransformCreate", ctx, input, model)}
}

func (_c *BaseCreateHandler_TransformCreate_Call) Run(run func(ctx context.Context, input interface{}, model db_repo.ModelBased)) *BaseCreateHandler_TransformCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(db_repo.ModelBased))
	})
	return _c
}

func (_c *BaseCreateHandler_TransformCreate_Call) Return(err error) *BaseCreateHandler_TransformCreate_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *BaseCreateHandler_TransformCreate_Call) RunAndReturn(run func(context.Context, interface{}, db_repo.ModelBased) error) *BaseCreateHandler_TransformCreate_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewBaseCreateHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBaseCreateHandler creates a new instance of BaseCreateHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBaseCreateHandler(t mockConstructorTestingTNewBaseCreateHandler) *BaseCreateHandler {
	mock := &BaseCreateHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}