// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"

	"github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	mock "github.com/stretchr/testify/mock"
)

// NewUserBackend creates a new instance of UserBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserBackend(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserBackend {
	mock := &UserBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// UserBackend is an autogenerated mock type for the UserBackend type
type UserBackend struct {
	mock.Mock
}

type UserBackend_Expecter struct {
	mock *mock.Mock
}

func (_m *UserBackend) EXPECT() *UserBackend_Expecter {
	return &UserBackend_Expecter{mock: &_m.Mock}
}

// Authenticate provides a mock function for the type UserBackend
func (_mock *UserBackend) Authenticate(ctx context.Context, username string, password string) (*userv1beta1.User, string, error) {
	ret := _mock.Called(ctx, username, password)

	if len(ret) == 0 {
		panic("no return value specified for Authenticate")
	}

	var r0 *userv1beta1.User
	var r1 string
	var r2 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) (*userv1beta1.User, string, error)); ok {
		return returnFunc(ctx, username, password)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) *userv1beta1.User); ok {
		r0 = returnFunc(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string) string); ok {
		r1 = returnFunc(ctx, username, password)
	} else {
		r1 = ret.Get(1).(string)
	}
	if returnFunc, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = returnFunc(ctx, username, password)
	} else {
		r2 = ret.Error(2)
	}
	return r0, r1, r2
}

// UserBackend_Authenticate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authenticate'
type UserBackend_Authenticate_Call struct {
	*mock.Call
}

// Authenticate is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - password string
func (_e *UserBackend_Expecter) Authenticate(ctx interface{}, username interface{}, password interface{}) *UserBackend_Authenticate_Call {
	return &UserBackend_Authenticate_Call{Call: _e.mock.On("Authenticate", ctx, username, password)}
}

func (_c *UserBackend_Authenticate_Call) Run(run func(ctx context.Context, username string, password string)) *UserBackend_Authenticate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *UserBackend_Authenticate_Call) Return(user *userv1beta1.User, s string, err error) *UserBackend_Authenticate_Call {
	_c.Call.Return(user, s, err)
	return _c
}

func (_c *UserBackend_Authenticate_Call) RunAndReturn(run func(ctx context.Context, username string, password string) (*userv1beta1.User, string, error)) *UserBackend_Authenticate_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUserFromClaims provides a mock function for the type UserBackend
func (_mock *UserBackend) CreateUserFromClaims(ctx context.Context, claims map[string]interface{}) (*userv1beta1.User, error) {
	ret := _mock.Called(ctx, claims)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserFromClaims")
	}

	var r0 *userv1beta1.User
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, map[string]interface{}) (*userv1beta1.User, error)); ok {
		return returnFunc(ctx, claims)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, map[string]interface{}) *userv1beta1.User); ok {
		r0 = returnFunc(ctx, claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = returnFunc(ctx, claims)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// UserBackend_CreateUserFromClaims_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserFromClaims'
type UserBackend_CreateUserFromClaims_Call struct {
	*mock.Call
}

// CreateUserFromClaims is a helper method to define mock.On call
//   - ctx context.Context
//   - claims map[string]interface{}
func (_e *UserBackend_Expecter) CreateUserFromClaims(ctx interface{}, claims interface{}) *UserBackend_CreateUserFromClaims_Call {
	return &UserBackend_CreateUserFromClaims_Call{Call: _e.mock.On("CreateUserFromClaims", ctx, claims)}
}

func (_c *UserBackend_CreateUserFromClaims_Call) Run(run func(ctx context.Context, claims map[string]interface{})) *UserBackend_CreateUserFromClaims_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 map[string]interface{}
		if args[1] != nil {
			arg1 = args[1].(map[string]interface{})
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *UserBackend_CreateUserFromClaims_Call) Return(user *userv1beta1.User, err error) *UserBackend_CreateUserFromClaims_Call {
	_c.Call.Return(user, err)
	return _c
}

func (_c *UserBackend_CreateUserFromClaims_Call) RunAndReturn(run func(ctx context.Context, claims map[string]interface{}) (*userv1beta1.User, error)) *UserBackend_CreateUserFromClaims_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByClaims provides a mock function for the type UserBackend
func (_mock *UserBackend) GetUserByClaims(ctx context.Context, claim string, value string) (*userv1beta1.User, string, error) {
	ret := _mock.Called(ctx, claim, value)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByClaims")
	}

	var r0 *userv1beta1.User
	var r1 string
	var r2 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) (*userv1beta1.User, string, error)); ok {
		return returnFunc(ctx, claim, value)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) *userv1beta1.User); ok {
		r0 = returnFunc(ctx, claim, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string) string); ok {
		r1 = returnFunc(ctx, claim, value)
	} else {
		r1 = ret.Get(1).(string)
	}
	if returnFunc, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = returnFunc(ctx, claim, value)
	} else {
		r2 = ret.Error(2)
	}
	return r0, r1, r2
}

// UserBackend_GetUserByClaims_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByClaims'
type UserBackend_GetUserByClaims_Call struct {
	*mock.Call
}

// GetUserByClaims is a helper method to define mock.On call
//   - ctx context.Context
//   - claim string
//   - value string
func (_e *UserBackend_Expecter) GetUserByClaims(ctx interface{}, claim interface{}, value interface{}) *UserBackend_GetUserByClaims_Call {
	return &UserBackend_GetUserByClaims_Call{Call: _e.mock.On("GetUserByClaims", ctx, claim, value)}
}

func (_c *UserBackend_GetUserByClaims_Call) Run(run func(ctx context.Context, claim string, value string)) *UserBackend_GetUserByClaims_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *UserBackend_GetUserByClaims_Call) Return(user *userv1beta1.User, s string, err error) *UserBackend_GetUserByClaims_Call {
	_c.Call.Return(user, s, err)
	return _c
}

func (_c *UserBackend_GetUserByClaims_Call) RunAndReturn(run func(ctx context.Context, claim string, value string) (*userv1beta1.User, string, error)) *UserBackend_GetUserByClaims_Call {
	_c.Call.Return(run)
	return _c
}

// SyncGroupMemberships provides a mock function for the type UserBackend
func (_mock *UserBackend) SyncGroupMemberships(ctx context.Context, user *userv1beta1.User, claims map[string]interface{}) error {
	ret := _mock.Called(ctx, user, claims)

	if len(ret) == 0 {
		panic("no return value specified for SyncGroupMemberships")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *userv1beta1.User, map[string]interface{}) error); ok {
		r0 = returnFunc(ctx, user, claims)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// UserBackend_SyncGroupMemberships_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncGroupMemberships'
type UserBackend_SyncGroupMemberships_Call struct {
	*mock.Call
}

// SyncGroupMemberships is a helper method to define mock.On call
//   - ctx context.Context
//   - user *userv1beta1.User
//   - claims map[string]interface{}
func (_e *UserBackend_Expecter) SyncGroupMemberships(ctx interface{}, user interface{}, claims interface{}) *UserBackend_SyncGroupMemberships_Call {
	return &UserBackend_SyncGroupMemberships_Call{Call: _e.mock.On("SyncGroupMemberships", ctx, user, claims)}
}

func (_c *UserBackend_SyncGroupMemberships_Call) Run(run func(ctx context.Context, user *userv1beta1.User, claims map[string]interface{})) *UserBackend_SyncGroupMemberships_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 *userv1beta1.User
		if args[1] != nil {
			arg1 = args[1].(*userv1beta1.User)
		}
		var arg2 map[string]interface{}
		if args[2] != nil {
			arg2 = args[2].(map[string]interface{})
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *UserBackend_SyncGroupMemberships_Call) Return(err error) *UserBackend_SyncGroupMemberships_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *UserBackend_SyncGroupMemberships_Call) RunAndReturn(run func(ctx context.Context, user *userv1beta1.User, claims map[string]interface{}) error) *UserBackend_SyncGroupMemberships_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserIfNeeded provides a mock function for the type UserBackend
func (_mock *UserBackend) UpdateUserIfNeeded(ctx context.Context, user *userv1beta1.User, claims map[string]interface{}) error {
	ret := _mock.Called(ctx, user, claims)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserIfNeeded")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *userv1beta1.User, map[string]interface{}) error); ok {
		r0 = returnFunc(ctx, user, claims)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// UserBackend_UpdateUserIfNeeded_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserIfNeeded'
type UserBackend_UpdateUserIfNeeded_Call struct {
	*mock.Call
}

// UpdateUserIfNeeded is a helper method to define mock.On call
//   - ctx context.Context
//   - user *userv1beta1.User
//   - claims map[string]interface{}
func (_e *UserBackend_Expecter) UpdateUserIfNeeded(ctx interface{}, user interface{}, claims interface{}) *UserBackend_UpdateUserIfNeeded_Call {
	return &UserBackend_UpdateUserIfNeeded_Call{Call: _e.mock.On("UpdateUserIfNeeded", ctx, user, claims)}
}

func (_c *UserBackend_UpdateUserIfNeeded_Call) Run(run func(ctx context.Context, user *userv1beta1.User, claims map[string]interface{})) *UserBackend_UpdateUserIfNeeded_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 *userv1beta1.User
		if args[1] != nil {
			arg1 = args[1].(*userv1beta1.User)
		}
		var arg2 map[string]interface{}
		if args[2] != nil {
			arg2 = args[2].(map[string]interface{})
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *UserBackend_UpdateUserIfNeeded_Call) Return(err error) *UserBackend_UpdateUserIfNeeded_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *UserBackend_UpdateUserIfNeeded_Call) RunAndReturn(run func(ctx context.Context, user *userv1beta1.User, claims map[string]interface{}) error) *UserBackend_UpdateUserIfNeeded_Call {
	_c.Call.Return(run)
	return _c
}
