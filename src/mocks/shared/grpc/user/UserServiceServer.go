// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	user "auth-svc/src/shared/grpc/user"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceServer is an autogenerated mock type for the UserServiceServer type
type UserServiceServer struct {
	mock.Mock
}

// Login provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) Login(_a0 context.Context, _a1 *user.LoginRequest) (*user.LoginResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *user.LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.LoginRequest) *user.LoginResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.LoginRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) Refresh(_a0 context.Context, _a1 *user.RefreshRequest) (*user.RefreshResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *user.RefreshResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.RefreshRequest) *user.RefreshResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.RefreshResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.RefreshRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) Register(_a0 context.Context, _a1 *user.RegisterRequest) (*user.RegisterResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *user.RegisterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.RegisterRequest) *user.RegisterResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.RegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.RegisterRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) User(_a0 context.Context, _a1 *user.UserRequest) (*user.UserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *user.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.UserRequest) *user.UserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.UserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedUserServiceServer provides a mock function with given fields:
func (_m *UserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewUserServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserServiceServer creates a new instance of UserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserServiceServer(t mockConstructorTestingTNewUserServiceServer) *UserServiceServer {
	mock := &UserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
