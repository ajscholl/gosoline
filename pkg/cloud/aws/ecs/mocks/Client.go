// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// DescribeContainerInstances provides a mock function with given fields: ctx, params, optFns
func (_m *Client) DescribeContainerInstances(ctx context.Context, params *ecs.DescribeContainerInstancesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeContainerInstances")
	}

	var r0 *ecs.DescribeContainerInstancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.DescribeContainerInstancesInput, ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.DescribeContainerInstancesInput, ...func(*ecs.Options)) *ecs.DescribeContainerInstancesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecs.DescribeContainerInstancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecs.DescribeContainerInstancesInput, ...func(*ecs.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DescribeContainerInstances_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeContainerInstances'
type Client_DescribeContainerInstances_Call struct {
	*mock.Call
}

// DescribeContainerInstances is a helper method to define mock.On call
//   - ctx context.Context
//   - params *ecs.DescribeContainerInstancesInput
//   - optFns ...func(*ecs.Options)
func (_e *Client_Expecter) DescribeContainerInstances(ctx interface{}, params interface{}, optFns ...interface{}) *Client_DescribeContainerInstances_Call {
	return &Client_DescribeContainerInstances_Call{Call: _e.mock.On("DescribeContainerInstances",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_DescribeContainerInstances_Call) Run(run func(ctx context.Context, params *ecs.DescribeContainerInstancesInput, optFns ...func(*ecs.Options))) *Client_DescribeContainerInstances_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*ecs.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*ecs.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*ecs.DescribeContainerInstancesInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_DescribeContainerInstances_Call) Return(_a0 *ecs.DescribeContainerInstancesOutput, _a1 error) *Client_DescribeContainerInstances_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DescribeContainerInstances_Call) RunAndReturn(run func(context.Context, *ecs.DescribeContainerInstancesInput, ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error)) *Client_DescribeContainerInstances_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeServices provides a mock function with given fields: ctx, params, optFns
func (_m *Client) DescribeServices(ctx context.Context, params *ecs.DescribeServicesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeServices")
	}

	var r0 *ecs.DescribeServicesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.DescribeServicesInput, ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.DescribeServicesInput, ...func(*ecs.Options)) *ecs.DescribeServicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecs.DescribeServicesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecs.DescribeServicesInput, ...func(*ecs.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DescribeServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeServices'
type Client_DescribeServices_Call struct {
	*mock.Call
}

// DescribeServices is a helper method to define mock.On call
//   - ctx context.Context
//   - params *ecs.DescribeServicesInput
//   - optFns ...func(*ecs.Options)
func (_e *Client_Expecter) DescribeServices(ctx interface{}, params interface{}, optFns ...interface{}) *Client_DescribeServices_Call {
	return &Client_DescribeServices_Call{Call: _e.mock.On("DescribeServices",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_DescribeServices_Call) Run(run func(ctx context.Context, params *ecs.DescribeServicesInput, optFns ...func(*ecs.Options))) *Client_DescribeServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*ecs.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*ecs.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*ecs.DescribeServicesInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_DescribeServices_Call) Return(_a0 *ecs.DescribeServicesOutput, _a1 error) *Client_DescribeServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DescribeServices_Call) RunAndReturn(run func(context.Context, *ecs.DescribeServicesInput, ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error)) *Client_DescribeServices_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeTasks provides a mock function with given fields: ctx, params, optFns
func (_m *Client) DescribeTasks(ctx context.Context, params *ecs.DescribeTasksInput, optFns ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeTasks")
	}

	var r0 *ecs.DescribeTasksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.DescribeTasksInput, ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.DescribeTasksInput, ...func(*ecs.Options)) *ecs.DescribeTasksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecs.DescribeTasksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecs.DescribeTasksInput, ...func(*ecs.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DescribeTasks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeTasks'
type Client_DescribeTasks_Call struct {
	*mock.Call
}

// DescribeTasks is a helper method to define mock.On call
//   - ctx context.Context
//   - params *ecs.DescribeTasksInput
//   - optFns ...func(*ecs.Options)
func (_e *Client_Expecter) DescribeTasks(ctx interface{}, params interface{}, optFns ...interface{}) *Client_DescribeTasks_Call {
	return &Client_DescribeTasks_Call{Call: _e.mock.On("DescribeTasks",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_DescribeTasks_Call) Run(run func(ctx context.Context, params *ecs.DescribeTasksInput, optFns ...func(*ecs.Options))) *Client_DescribeTasks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*ecs.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*ecs.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*ecs.DescribeTasksInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_DescribeTasks_Call) Return(_a0 *ecs.DescribeTasksOutput, _a1 error) *Client_DescribeTasks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DescribeTasks_Call) RunAndReturn(run func(context.Context, *ecs.DescribeTasksInput, ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error)) *Client_DescribeTasks_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateService provides a mock function with given fields: ctx, params, optFns
func (_m *Client) UpdateService(ctx context.Context, params *ecs.UpdateServiceInput, optFns ...func(*ecs.Options)) (*ecs.UpdateServiceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateService")
	}

	var r0 *ecs.UpdateServiceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.UpdateServiceInput, ...func(*ecs.Options)) (*ecs.UpdateServiceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecs.UpdateServiceInput, ...func(*ecs.Options)) *ecs.UpdateServiceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecs.UpdateServiceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecs.UpdateServiceInput, ...func(*ecs.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_UpdateService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateService'
type Client_UpdateService_Call struct {
	*mock.Call
}

// UpdateService is a helper method to define mock.On call
//   - ctx context.Context
//   - params *ecs.UpdateServiceInput
//   - optFns ...func(*ecs.Options)
func (_e *Client_Expecter) UpdateService(ctx interface{}, params interface{}, optFns ...interface{}) *Client_UpdateService_Call {
	return &Client_UpdateService_Call{Call: _e.mock.On("UpdateService",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_UpdateService_Call) Run(run func(ctx context.Context, params *ecs.UpdateServiceInput, optFns ...func(*ecs.Options))) *Client_UpdateService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*ecs.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*ecs.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*ecs.UpdateServiceInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_UpdateService_Call) Return(_a0 *ecs.UpdateServiceOutput, _a1 error) *Client_UpdateService_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_UpdateService_Call) RunAndReturn(run func(context.Context, *ecs.UpdateServiceInput, ...func(*ecs.Options)) (*ecs.UpdateServiceOutput, error)) *Client_UpdateService_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
