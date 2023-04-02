// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mongo "go_be_clean_architecture/mongo"
	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *Database) Client() mongo.Client {
	ret := _m.Called()

	var r0 mongo.Client
	if rf, ok := ret.Get(0).(func() mongo.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Client)
		}
	}

	return r0
}

// Collection provides a mock function with given fields: _a0
func (_m *Database) Collection(_a0 string) mongo.Collection {
	ret := _m.Called(_a0)

	var r0 mongo.Collection
	if rf, ok := ret.Get(0).(func(string) mongo.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Collection)
		}
	}

	return r0
}

type mockConstructorTestingTNewDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabase(t mockConstructorTestingTNewDatabase) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
