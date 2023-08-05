// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mdl "github.com/justtrackio/gosoline/pkg/mdl"

	mock "github.com/stretchr/testify/mock"
)

// BaseCreateHandler is an autogenerated mock type for the BaseCreateHandler type
type BaseCreateHandler[I interface{}, K mdl.PossibleIdentifier, M db_repo.ModelBased[K]] struct {
	mock.Mock
}

type BaseCreateHandler_Expecter[I interface{}, K mdl.PossibleIdentifier, M db_repo.ModelBased[K]] struct {
	mock *mock.Mock
}

func (_m *BaseCreateHandler[I, K, M]) EXPECT() *BaseCreateHandler_Expecter[I, K, M] {
	return &BaseCreateHandler_Expecter[I, K, M]{mock: &_m.Mock}
}

// TransformCreate provides a mock function with given fields: ctx, input
func (_m *BaseCreateHandler[I, K, M]) TransformCreate(ctx context.Context, input *I) (M, error) {
	ret := _m.Called(ctx, input)

	var r0 M
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *I) (M, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *I) M); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(M)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *I) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BaseCreateHandler_TransformCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransformCreate'
type BaseCreateHandler_TransformCreate_Call[I interface{}, K mdl.PossibleIdentifier, M db_repo.ModelBased[K]] struct {
	*mock.Call
}

// TransformCreate is a helper method to define mock.On call
//   - ctx context.Context
//   - input *I
func (_e *BaseCreateHandler_Expecter[I, K, M]) TransformCreate(ctx interface{}, input interface{}) *BaseCreateHandler_TransformCreate_Call[I, K, M] {
	return &BaseCreateHandler_TransformCreate_Call[I, K, M]{Call: _e.mock.On("TransformCreate", ctx, input)}
}

func (_c *BaseCreateHandler_TransformCreate_Call[I, K, M]) Run(run func(ctx context.Context, input *I)) *BaseCreateHandler_TransformCreate_Call[I, K, M] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*I))
	})
	return _c
}

func (_c *BaseCreateHandler_TransformCreate_Call[I, K, M]) Return(_a0 M, _a1 error) *BaseCreateHandler_TransformCreate_Call[I, K, M] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BaseCreateHandler_TransformCreate_Call[I, K, M]) RunAndReturn(run func(context.Context, *I) (M, error)) *BaseCreateHandler_TransformCreate_Call[I, K, M] {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewBaseCreateHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBaseCreateHandler creates a new instance of BaseCreateHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBaseCreateHandler[I interface{}, K mdl.PossibleIdentifier, M db_repo.ModelBased[K]](t mockConstructorTestingTNewBaseCreateHandler) *BaseCreateHandler[I, K, M] {
	mock := &BaseCreateHandler[I, K, M]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
