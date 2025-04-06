// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/So-ham/Content-Moderation/internal/entities"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: ctx, req
func (_m *Service) AddComment(ctx context.Context, req *entities.CommentRequest) (*entities.CommentResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for AddComment")
	}

	var r0 *entities.CommentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.CommentRequest) (*entities.CommentResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.CommentRequest) *entities.CommentResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.CommentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.CommentRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddReview provides a mock function with given fields: ctx, req
func (_m *Service) AddReview(ctx context.Context, req *entities.ReviewRequest) (*entities.ReviewResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for AddReview")
	}

	var r0 *entities.ReviewResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ReviewRequest) (*entities.ReviewResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ReviewRequest) *entities.ReviewResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ReviewResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.ReviewRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPosts provides a mock function with no fields
func (_m *Service) GetAllPosts() ([]entities.PostResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllPosts")
	}

	var r0 []entities.PostResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entities.PostResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entities.PostResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.PostResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: req
func (_m *Service) Login(req *entities.UserLoginRequest) (*entities.UserResponse, string, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *entities.UserResponse
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*entities.UserLoginRequest) (*entities.UserResponse, string, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*entities.UserLoginRequest) *entities.UserResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.UserLoginRequest) string); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*entities.UserLoginRequest) error); ok {
		r2 = rf(req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Signup provides a mock function with given fields: ctx, req
func (_m *Service) Signup(ctx context.Context, req *entities.UserSignupRequest) (*entities.UserResponse, string, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Signup")
	}

	var r0 *entities.UserResponse
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserSignupRequest) (*entities.UserResponse, string, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserSignupRequest) *entities.UserResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.UserSignupRequest) string); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *entities.UserSignupRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
