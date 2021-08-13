// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	ddb "github.com/applike/gosoline/pkg/ddb"
	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	expression "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

	mock "github.com/stretchr/testify/mock"
)

// UpdateItemBuilder is an autogenerated mock type for the UpdateItemBuilder type
type UpdateItemBuilder struct {
	mock.Mock
}

// Add provides a mock function with given fields: path, value
func (_m *UpdateItemBuilder) Add(path string, value interface{}) ddb.UpdateItemBuilder {
	ret := _m.Called(path, value)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(string, interface{}) ddb.UpdateItemBuilder); ok {
		r0 = rf(path, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// Build provides a mock function with given fields: item
func (_m *UpdateItemBuilder) Build(item interface{}) (*dynamodb.UpdateItemInput, error) {
	ret := _m.Called(item)

	var r0 *dynamodb.UpdateItemInput
	if rf, ok := ret.Get(0).(func(interface{}) *dynamodb.UpdateItemInput); ok {
		r0 = rf(item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.UpdateItemInput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: path, value
func (_m *UpdateItemBuilder) Delete(path string, value interface{}) ddb.UpdateItemBuilder {
	ret := _m.Called(path, value)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(string, interface{}) ddb.UpdateItemBuilder); ok {
		r0 = rf(path, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: path
func (_m *UpdateItemBuilder) Remove(path string) ddb.UpdateItemBuilder {
	ret := _m.Called(path)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(string) ddb.UpdateItemBuilder); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// RemoveMultiple provides a mock function with given fields: paths
func (_m *UpdateItemBuilder) RemoveMultiple(paths ...string) ddb.UpdateItemBuilder {
	_va := make([]interface{}, len(paths))
	for _i := range paths {
		_va[_i] = paths[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(...string) ddb.UpdateItemBuilder); ok {
		r0 = rf(paths...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// ReturnAllNew provides a mock function with given fields:
func (_m *UpdateItemBuilder) ReturnAllNew() ddb.UpdateItemBuilder {
	ret := _m.Called()

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.UpdateItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// ReturnAllOld provides a mock function with given fields:
func (_m *UpdateItemBuilder) ReturnAllOld() ddb.UpdateItemBuilder {
	ret := _m.Called()

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.UpdateItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// ReturnNone provides a mock function with given fields:
func (_m *UpdateItemBuilder) ReturnNone() ddb.UpdateItemBuilder {
	ret := _m.Called()

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.UpdateItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// ReturnUpdatedNew provides a mock function with given fields:
func (_m *UpdateItemBuilder) ReturnUpdatedNew() ddb.UpdateItemBuilder {
	ret := _m.Called()

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.UpdateItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// ReturnUpdatedOld provides a mock function with given fields:
func (_m *UpdateItemBuilder) ReturnUpdatedOld() ddb.UpdateItemBuilder {
	ret := _m.Called()

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.UpdateItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// Set provides a mock function with given fields: path, value
func (_m *UpdateItemBuilder) Set(path string, value interface{}) ddb.UpdateItemBuilder {
	ret := _m.Called(path, value)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(string, interface{}) ddb.UpdateItemBuilder); ok {
		r0 = rf(path, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// SetIfNotExist provides a mock function with given fields: path, value
func (_m *UpdateItemBuilder) SetIfNotExist(path string, value interface{}) ddb.UpdateItemBuilder {
	ret := _m.Called(path, value)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(string, interface{}) ddb.UpdateItemBuilder); ok {
		r0 = rf(path, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// SetMap provides a mock function with given fields: values
func (_m *UpdateItemBuilder) SetMap(values map[string]interface{}) ddb.UpdateItemBuilder {
	ret := _m.Called(values)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(map[string]interface{}) ddb.UpdateItemBuilder); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// WithCondition provides a mock function with given fields: cond
func (_m *UpdateItemBuilder) WithCondition(cond expression.ConditionBuilder) ddb.UpdateItemBuilder {
	ret := _m.Called(cond)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(expression.ConditionBuilder) ddb.UpdateItemBuilder); ok {
		r0 = rf(cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// WithHash provides a mock function with given fields: hashValue
func (_m *UpdateItemBuilder) WithHash(hashValue interface{}) ddb.UpdateItemBuilder {
	ret := _m.Called(hashValue)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.UpdateItemBuilder); ok {
		r0 = rf(hashValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// WithRange provides a mock function with given fields: rangeValue
func (_m *UpdateItemBuilder) WithRange(rangeValue interface{}) ddb.UpdateItemBuilder {
	ret := _m.Called(rangeValue)

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.UpdateItemBuilder); ok {
		r0 = rf(rangeValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}
