// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import kernel "github.com/applike/gosoline/pkg/kernel"
import mock "github.com/stretchr/testify/mock"

// Kernel is an autogenerated mock type for the Kernel type
type Kernel struct {
	mock.Mock
}

// Add provides a mock function with given fields: name, module
func (_m *Kernel) Add(name string, module kernel.Module) {
	_m.Called(name, module)
}

// AddFactory provides a mock function with given fields: factory
func (_m *Kernel) AddFactory(factory kernel.ModuleFactory) {
	_m.Called(factory)
}

// Booted provides a mock function with given fields:
func (_m *Kernel) Booted() <-chan struct{} {
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

// Run provides a mock function with given fields:
func (_m *Kernel) Run() {
	_m.Called()
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

// Stop provides a mock function with given fields: reason
func (_m *Kernel) Stop(reason string) {
	_m.Called(reason)
}
