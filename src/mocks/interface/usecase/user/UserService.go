// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "auth-svc/src/shared/grpc/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserService) FindByEmail(ctx context.Context, email string) (*user.UserResponse, error) {
	ret := _m.Called(ctx, email)

	var r0 *user.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *user.UserResponse); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, request
func (_m *UserService) FindByID(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *user.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.UserRequest) *user.UserResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.UserRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, request
func (_m *UserService) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *user.LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.LoginRequest) *user.LoginResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.LoginRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: ctx, request
func (_m *UserService) Refresh(ctx context.Context, request *user.RefreshRequest) (*user.RefreshResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *user.RefreshResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.RefreshRequest) *user.RefreshResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.RefreshResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.RefreshRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, request
func (_m *UserService) Register(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *user.RegisterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.RegisterRequest) *user.RegisterResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.RegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.RegisterRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}