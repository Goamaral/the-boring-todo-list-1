// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gorm "gorm.io/gorm"

	gormprovider "example.com/the-boring-to-do-list-1/pkg/provider/gorm"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// NewQuery provides a mock function with given fields: ctx
func (_m *Repository) NewQuery(ctx context.Context) *gorm.DB {
	ret := _m.Called(ctx)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(context.Context) *gorm.DB); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// NewQueryWithOpts provides a mock function with given fields: ctx, opts
func (_m *Repository) NewQueryWithOpts(ctx context.Context, opts ...gormprovider.QueryOption) *gorm.DB {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(context.Context, ...gormprovider.QueryOption) *gorm.DB); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// TableName provides a mock function with given fields:
func (_m *Repository) TableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
