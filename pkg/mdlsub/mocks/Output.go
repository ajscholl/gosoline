// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mdlsub "github.com/justtrackio/gosoline/pkg/mdlsub"
	mock "github.com/stretchr/testify/mock"
)

// Output is an autogenerated mock type for the Output type
type Output struct {
	mock.Mock
}

type Output_Expecter struct {
	mock *mock.Mock
}

func (_m *Output) EXPECT() *Output_Expecter {
	return &Output_Expecter{mock: &_m.Mock}
}

// Persist provides a mock function with given fields: ctx, model, op
func (_m *Output) Persist(ctx context.Context, model mdlsub.Model, op string) error {
	ret := _m.Called(ctx, model, op)

	if len(ret) == 0 {
		panic("no return value specified for Persist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mdlsub.Model, string) error); ok {
		r0 = rf(ctx, model, op)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Output_Persist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Persist'
type Output_Persist_Call struct {
	*mock.Call
}

// Persist is a helper method to define mock.On call
//   - ctx context.Context
//   - model mdlsub.Model
//   - op string
func (_e *Output_Expecter) Persist(ctx interface{}, model interface{}, op interface{}) *Output_Persist_Call {
	return &Output_Persist_Call{Call: _e.mock.On("Persist", ctx, model, op)}
}

func (_c *Output_Persist_Call) Run(run func(ctx context.Context, model mdlsub.Model, op string)) *Output_Persist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(mdlsub.Model), args[2].(string))
	})
	return _c
}

func (_c *Output_Persist_Call) Return(_a0 error) *Output_Persist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Output_Persist_Call) RunAndReturn(run func(context.Context, mdlsub.Model, string) error) *Output_Persist_Call {
	_c.Call.Return(run)
	return _c
}

// NewOutput creates a new instance of Output. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOutput(t interface {
	mock.TestingT
	Cleanup(func())
}) *Output {
	mock := &Output{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
