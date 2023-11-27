// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	medicinecategories "MiniProject/features/medicine_categories"

	mock "github.com/stretchr/testify/mock"
)

// MedicineCategoryDataInterface is an autogenerated mock type for the MedicineCategoryDataInterface type
type MedicineCategoryDataInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *MedicineCategoryDataInterface) Delete(id int) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *MedicineCategoryDataInterface) GetAll() ([]medicinecategories.MedicineCategory, error) {
	ret := _m.Called()

	var r0 []medicinecategories.MedicineCategory
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]medicinecategories.MedicineCategory, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []medicinecategories.MedicineCategory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]medicinecategories.MedicineCategory)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *MedicineCategoryDataInterface) GetByID(id int) ([]medicinecategories.MedicineCategory, error) {
	ret := _m.Called(id)

	var r0 []medicinecategories.MedicineCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]medicinecategories.MedicineCategory, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []medicinecategories.MedicineCategory); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]medicinecategories.MedicineCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newData
func (_m *MedicineCategoryDataInterface) Insert(newData medicinecategories.MedicineCategory) (*medicinecategories.MedicineCategory, error) {
	ret := _m.Called(newData)

	var r0 *medicinecategories.MedicineCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(medicinecategories.MedicineCategory) (*medicinecategories.MedicineCategory, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(medicinecategories.MedicineCategory) *medicinecategories.MedicineCategory); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*medicinecategories.MedicineCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(medicinecategories.MedicineCategory) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: newData, id
func (_m *MedicineCategoryDataInterface) Update(newData medicinecategories.UpdateMedicineCategory, id int) (bool, error) {
	ret := _m.Called(newData, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(medicinecategories.UpdateMedicineCategory, int) (bool, error)); ok {
		return rf(newData, id)
	}
	if rf, ok := ret.Get(0).(func(medicinecategories.UpdateMedicineCategory, int) bool); ok {
		r0 = rf(newData, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(medicinecategories.UpdateMedicineCategory, int) error); ok {
		r1 = rf(newData, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMedicineCategoryDataInterface creates a new instance of MedicineCategoryDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMedicineCategoryDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MedicineCategoryDataInterface {
	mock := &MedicineCategoryDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}