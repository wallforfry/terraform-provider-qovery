// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks_test

import (
	context "context"

	secret "github.com/qovery/terraform-provider-qovery/internal/domain/secret"
	mock "github.com/stretchr/testify/mock"
)

// SecretRepository is an autogenerated mock type for the Repository type
type SecretRepository struct {
	mock.Mock
}

type SecretRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *SecretRepository) EXPECT() *SecretRepository_Expecter {
	return &SecretRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, scopeResourceID, request
func (_m *SecretRepository) Create(ctx context.Context, scopeResourceID string, request secret.UpsertRequest) (*secret.Secret, error) {
	ret := _m.Called(ctx, scopeResourceID, request)

	var r0 *secret.Secret
	if rf, ok := ret.Get(0).(func(context.Context, string, secret.UpsertRequest) *secret.Secret); ok {
		r0 = rf(ctx, scopeResourceID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secret.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, secret.UpsertRequest) error); ok {
		r1 = rf(ctx, scopeResourceID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SecretRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type SecretRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
//   - request secret.UpsertRequest
func (_e *SecretRepository_Expecter) Create(ctx interface{}, scopeResourceID interface{}, request interface{}) *SecretRepository_Create_Call {
	return &SecretRepository_Create_Call{Call: _e.mock.On("Create", ctx, scopeResourceID, request)}
}

func (_c *SecretRepository_Create_Call) Run(run func(ctx context.Context, scopeResourceID string, request secret.UpsertRequest)) *SecretRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(secret.UpsertRequest))
	})
	return _c
}

func (_c *SecretRepository_Create_Call) Return(_a0 *secret.Secret, _a1 error) *SecretRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Delete provides a mock function with given fields: ctx, scopeResourceID, secretID
func (_m *SecretRepository) Delete(ctx context.Context, scopeResourceID string, secretID string) error {
	ret := _m.Called(ctx, scopeResourceID, secretID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, scopeResourceID, secretID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SecretRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type SecretRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
//   - secretID string
func (_e *SecretRepository_Expecter) Delete(ctx interface{}, scopeResourceID interface{}, secretID interface{}) *SecretRepository_Delete_Call {
	return &SecretRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, scopeResourceID, secretID)}
}

func (_c *SecretRepository_Delete_Call) Run(run func(ctx context.Context, scopeResourceID string, secretID string)) *SecretRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *SecretRepository_Delete_Call) Return(_a0 error) *SecretRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// List provides a mock function with given fields: ctx, scopeResourceID
func (_m *SecretRepository) List(ctx context.Context, scopeResourceID string) (secret.Secrets, error) {
	ret := _m.Called(ctx, scopeResourceID)

	var r0 secret.Secrets
	if rf, ok := ret.Get(0).(func(context.Context, string) secret.Secrets); ok {
		r0 = rf(ctx, scopeResourceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(secret.Secrets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, scopeResourceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SecretRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type SecretRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
func (_e *SecretRepository_Expecter) List(ctx interface{}, scopeResourceID interface{}) *SecretRepository_List_Call {
	return &SecretRepository_List_Call{Call: _e.mock.On("List", ctx, scopeResourceID)}
}

func (_c *SecretRepository_List_Call) Run(run func(ctx context.Context, scopeResourceID string)) *SecretRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *SecretRepository_List_Call) Return(_a0 secret.Secrets, _a1 error) *SecretRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: ctx, scopeResourceID, secretID, request
func (_m *SecretRepository) Update(ctx context.Context, scopeResourceID string, secretID string, request secret.UpsertRequest) (*secret.Secret, error) {
	ret := _m.Called(ctx, scopeResourceID, secretID, request)

	var r0 *secret.Secret
	if rf, ok := ret.Get(0).(func(context.Context, string, string, secret.UpsertRequest) *secret.Secret); ok {
		r0 = rf(ctx, scopeResourceID, secretID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secret.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, secret.UpsertRequest) error); ok {
		r1 = rf(ctx, scopeResourceID, secretID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SecretRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type SecretRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
//   - secretID string
//   - request secret.UpsertRequest
func (_e *SecretRepository_Expecter) Update(ctx interface{}, scopeResourceID interface{}, secretID interface{}, request interface{}) *SecretRepository_Update_Call {
	return &SecretRepository_Update_Call{Call: _e.mock.On("Update", ctx, scopeResourceID, secretID, request)}
}

func (_c *SecretRepository_Update_Call) Run(run func(ctx context.Context, scopeResourceID string, secretID string, request secret.UpsertRequest)) *SecretRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(secret.UpsertRequest))
	})
	return _c
}

func (_c *SecretRepository_Update_Call) Return(_a0 *secret.Secret, _a1 error) *SecretRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewSecretRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewSecretRepository creates a new instance of SecretRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSecretRepository(t mockConstructorTestingTNewSecretRepository) *SecretRepository {
	mock := &SecretRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
