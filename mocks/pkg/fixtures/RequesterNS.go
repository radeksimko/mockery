// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequesterNS is an autogenerated mock type for the RequesterNS type
type RequesterNS struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterNS) Get(path string) (http.Response, error) {
	ret := _m.Called(path)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string) http.Response); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(http.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRequesterNST interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequesterNS creates a new instance of RequesterNS. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterNS(t NewRequesterNST) *RequesterNS {
	mock := &RequesterNS{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
