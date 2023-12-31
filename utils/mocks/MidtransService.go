// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	transactions "MiniProject/features/transactions"

	mock "github.com/stretchr/testify/mock"
)

// MidtransService is an autogenerated mock type for the MidtransService type
type MidtransService struct {
	mock.Mock
}

// GenerateSnapURL provides a mock function with given fields: _a0
func (_m *MidtransService) GenerateSnapURL(_a0 transactions.Transaction) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(transactions.Transaction) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(transactions.Transaction) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(transactions.Transaction) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotification provides a mock function with given fields: transID
func (_m *MidtransService) GetNotification(transID string) (string, error) {
	ret := _m.Called(transID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(transID)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(transID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(transID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMidtransService creates a new instance of MidtransService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMidtransService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MidtransService {
	mock := &MidtransService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
