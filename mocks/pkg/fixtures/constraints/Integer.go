// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Integer is an autogenerated mock type for the Integer type
type Integer struct {
	mock.Mock
}

type NewIntegerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewInteger creates a new instance of Integer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInteger(t NewIntegerT) *Integer {
	mock := &Integer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
