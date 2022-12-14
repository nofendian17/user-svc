// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"
)

// DatabaseClient is an autogenerated mock type for the DatabaseClient type
type DatabaseClient struct {
	mock.Mock
}

// Conn provides a mock function with given fields:
func (_m *DatabaseClient) Conn() (*sqlx.DB, error) {
	ret := _m.Called()

	var r0 *sqlx.DB
	if rf, ok := ret.Get(0).(func() *sqlx.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.DB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disconnect provides a mock function with given fields:
func (_m *DatabaseClient) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: ctx
func (_m *DatabaseClient) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDatabaseClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabaseClient creates a new instance of DatabaseClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabaseClient(t mockConstructorTestingTNewDatabaseClient) *DatabaseClient {
	mock := &DatabaseClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
