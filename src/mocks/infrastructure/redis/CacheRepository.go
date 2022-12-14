// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	redis "auth-svc/src/infrastructure/redis"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CacheRepository is an autogenerated mock type for the CacheRepository type
type CacheRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, key
func (_m *CacheRepository) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key
func (_m *CacheRepository) Get(ctx context.Context, key string) (*redis.CacheContainer, error) {
	ret := _m.Called(ctx, key)

	var r0 *redis.CacheContainer
	if rf, ok := ret.Get(0).(func(context.Context, string) *redis.CacheContainer); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.CacheContainer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, key, cache, ttl
func (_m *CacheRepository) Set(ctx context.Context, key string, cache redis.CacheContainer, ttl int64) error {
	ret := _m.Called(ctx, key, cache, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, redis.CacheContainer, int64) error); ok {
		r0 = rf(ctx, key, cache, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCacheRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheRepository creates a new instance of CacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheRepository(t mockConstructorTestingTNewCacheRepository) *CacheRepository {
	mock := &CacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
