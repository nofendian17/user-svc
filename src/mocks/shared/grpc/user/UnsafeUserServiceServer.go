// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeUserServiceServer is an autogenerated mock type for the UnsafeUserServiceServer type
type UnsafeUserServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedUserServiceServer provides a mock function with given fields:
func (_m *UnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeUserServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeUserServiceServer creates a new instance of UnsafeUserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeUserServiceServer(t mockConstructorTestingTNewUnsafeUserServiceServer) *UnsafeUserServiceServer {
	mock := &UnsafeUserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
