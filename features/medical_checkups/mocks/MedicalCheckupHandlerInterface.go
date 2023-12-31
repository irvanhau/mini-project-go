// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// MedicalCheckupHandlerInterface is an autogenerated mock type for the MedicalCheckupHandlerInterface type
type MedicalCheckupHandlerInterface struct {
	mock.Mock
}

// CreateMedicalCheckup provides a mock function with given fields:
func (_m *MedicalCheckupHandlerInterface) CreateMedicalCheckup() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// DeleteMedicalCheckup provides a mock function with given fields:
func (_m *MedicalCheckupHandlerInterface) DeleteMedicalCheckup() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetMedicalCheckup provides a mock function with given fields:
func (_m *MedicalCheckupHandlerInterface) GetMedicalCheckup() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetMedicalCheckups provides a mock function with given fields:
func (_m *MedicalCheckupHandlerInterface) GetMedicalCheckups() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// UpdateMedicalCheckup provides a mock function with given fields:
func (_m *MedicalCheckupHandlerInterface) UpdateMedicalCheckup() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewMedicalCheckupHandlerInterface creates a new instance of MedicalCheckupHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMedicalCheckupHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MedicalCheckupHandlerInterface {
	mock := &MedicalCheckupHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
