// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain "project2/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// ClientRepo is an autogenerated mock type for the clientRepo type
type ClientRepo struct {
	mock.Mock
}

type ClientRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientRepo) EXPECT() *ClientRepo_Expecter {
	return &ClientRepo_Expecter{mock: &_m.Mock}
}

// FindByID provides a mock function with given fields: id
func (_m *ClientRepo) FindByID(id int) (*domain.Client, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *domain.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.Client, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Client); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRepo_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type ClientRepo_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - id int
func (_e *ClientRepo_Expecter) FindByID(id interface{}) *ClientRepo_FindByID_Call {
	return &ClientRepo_FindByID_Call{Call: _e.mock.On("FindByID", id)}
}

func (_c *ClientRepo_FindByID_Call) Run(run func(id int)) *ClientRepo_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ClientRepo_FindByID_Call) Return(_a0 *domain.Client, _a1 error) *ClientRepo_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientRepo_FindByID_Call) RunAndReturn(run func(int) (*domain.Client, error)) *ClientRepo_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientRepo creates a new instance of ClientRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientRepo {
	mock := &ClientRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
