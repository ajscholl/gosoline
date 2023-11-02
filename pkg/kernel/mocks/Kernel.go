// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	kernel "github.com/justtrackio/gosoline/pkg/kernel"
	mock "github.com/stretchr/testify/mock"
)

// Kernel is an autogenerated mock type for the Kernel type
type Kernel struct {
	mock.Mock
}

type Kernel_Expecter struct {
	mock *mock.Mock
}

func (_m *Kernel) EXPECT() *Kernel_Expecter {
	return &Kernel_Expecter{mock: &_m.Mock}
}

// HealthCheck provides a mock function with given fields:
func (_m *Kernel) HealthCheck() kernel.HealthCheckResult {
	ret := _m.Called()

	var r0 kernel.HealthCheckResult
	if rf, ok := ret.Get(0).(func() kernel.HealthCheckResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kernel.HealthCheckResult)
		}
	}

	return r0
}

// Kernel_HealthCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthCheck'
type Kernel_HealthCheck_Call struct {
	*mock.Call
}

// HealthCheck is a helper method to define mock.On call
func (_e *Kernel_Expecter) HealthCheck() *Kernel_HealthCheck_Call {
	return &Kernel_HealthCheck_Call{Call: _e.mock.On("HealthCheck")}
}

func (_c *Kernel_HealthCheck_Call) Run(run func()) *Kernel_HealthCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Kernel_HealthCheck_Call) Return(_a0 kernel.HealthCheckResult) *Kernel_HealthCheck_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Kernel_HealthCheck_Call) RunAndReturn(run func() kernel.HealthCheckResult) *Kernel_HealthCheck_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields:
func (_m *Kernel) Run() {
	_m.Called()
}

// Kernel_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type Kernel_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *Kernel_Expecter) Run() *Kernel_Run_Call {
	return &Kernel_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *Kernel_Run_Call) Run(run func()) *Kernel_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Kernel_Run_Call) Return() *Kernel_Run_Call {
	_c.Call.Return()
	return _c
}

func (_c *Kernel_Run_Call) RunAndReturn(run func()) *Kernel_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Running provides a mock function with given fields:
func (_m *Kernel) Running() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Kernel_Running_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Running'
type Kernel_Running_Call struct {
	*mock.Call
}

// Running is a helper method to define mock.On call
func (_e *Kernel_Expecter) Running() *Kernel_Running_Call {
	return &Kernel_Running_Call{Call: _e.mock.On("Running")}
}

func (_c *Kernel_Running_Call) Run(run func()) *Kernel_Running_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Kernel_Running_Call) Return(_a0 <-chan struct{}) *Kernel_Running_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Kernel_Running_Call) RunAndReturn(run func() <-chan struct{}) *Kernel_Running_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: reason
func (_m *Kernel) Stop(reason string) {
	_m.Called(reason)
}

// Kernel_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Kernel_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - reason string
func (_e *Kernel_Expecter) Stop(reason interface{}) *Kernel_Stop_Call {
	return &Kernel_Stop_Call{Call: _e.mock.On("Stop", reason)}
}

func (_c *Kernel_Stop_Call) Run(run func(reason string)) *Kernel_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Kernel_Stop_Call) Return() *Kernel_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *Kernel_Stop_Call) RunAndReturn(run func(string)) *Kernel_Stop_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewKernel interface {
	mock.TestingT
	Cleanup(func())
}

// NewKernel creates a new instance of Kernel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKernel(t mockConstructorTestingTNewKernel) *Kernel {
	mock := &Kernel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
