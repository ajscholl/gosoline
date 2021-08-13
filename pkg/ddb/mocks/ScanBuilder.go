// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	ddb "github.com/applike/gosoline/pkg/ddb"
	expression "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	mock "github.com/stretchr/testify/mock"
)

// ScanBuilder is an autogenerated mock type for the ScanBuilder type
type ScanBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields: result
func (_m *ScanBuilder) Build(result interface{}) (*ddb.ScanOperation, error) {
	ret := _m.Called(result)

	var r0 *ddb.ScanOperation
	if rf, ok := ret.Get(0).(func(interface{}) *ddb.ScanOperation); ok {
		r0 = rf(result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.ScanOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableTtlFilter provides a mock function with given fields:
func (_m *ScanBuilder) DisableTtlFilter() ddb.ScanBuilder {
	ret := _m.Called()

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func() ddb.ScanBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// WithConsistentRead provides a mock function with given fields: consistentRead
func (_m *ScanBuilder) WithConsistentRead(consistentRead bool) ddb.ScanBuilder {
	ret := _m.Called(consistentRead)

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func(bool) ddb.ScanBuilder); ok {
		r0 = rf(consistentRead)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// WithFilter provides a mock function with given fields: filter
func (_m *ScanBuilder) WithFilter(filter expression.ConditionBuilder) ddb.ScanBuilder {
	ret := _m.Called(filter)

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func(expression.ConditionBuilder) ddb.ScanBuilder); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// WithIndex provides a mock function with given fields: name
func (_m *ScanBuilder) WithIndex(name string) ddb.ScanBuilder {
	ret := _m.Called(name)

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func(string) ddb.ScanBuilder); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// WithLimit provides a mock function with given fields: limit
func (_m *ScanBuilder) WithLimit(limit int) ddb.ScanBuilder {
	ret := _m.Called(limit)

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func(int) ddb.ScanBuilder); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// WithPageSize provides a mock function with given fields: size
func (_m *ScanBuilder) WithPageSize(size int) ddb.ScanBuilder {
	ret := _m.Called(size)

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func(int) ddb.ScanBuilder); ok {
		r0 = rf(size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// WithProjection provides a mock function with given fields: projection
func (_m *ScanBuilder) WithProjection(projection interface{}) ddb.ScanBuilder {
	ret := _m.Called(projection)

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.ScanBuilder); ok {
		r0 = rf(projection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// WithSegment provides a mock function with given fields: segment, total
func (_m *ScanBuilder) WithSegment(segment int, total int) ddb.ScanBuilder {
	ret := _m.Called(segment, total)

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func(int, int) ddb.ScanBuilder); ok {
		r0 = rf(segment, total)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}
