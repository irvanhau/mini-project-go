package service

import (
	medicinecategories "MiniProject/features/medicine_categories"
	"MiniProject/features/medicine_categories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMedicineCategories(t *testing.T) {
	data := mocks.NewMedicineCategoryDataInterface(t)
	service := New(data)
	medicineCategories := []medicinecategories.MedicineCategory{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(medicineCategories, nil).Once()

		result, err := service.GetMedicineCategories()
		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		result, err := service.GetMedicineCategories()
		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Process Failed")
		assert.Nil(t, result)
	})
}

func TestGetMedicineCategory(t *testing.T) {
	data := mocks.NewMedicineCategoryDataInterface(t)
	service := New(data)
	medicineCategories := []medicinecategories.MedicineCategory{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(medicineCategories, nil).Once()

		result, err := service.GetMedicineCategory(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Failed")).Once()

		result, err := service.GetMedicineCategory(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "Get By ID Failed")
		assert.Nil(t, result)
	})
}

func TestCreateMedicineCategory(t *testing.T) {
	data := mocks.NewMedicineCategoryDataInterface(t)
	service := New(data)
	medicineCategory := medicinecategories.MedicineCategory{
		Name:        "Panadol",
		Description: "Panadol",
	}

	t.Run("Success Insert", func(t *testing.T) {
		data.On("Insert", medicineCategory).Return(&medicineCategory, nil).Once()

		result, err := service.CreateMedicineCategory(medicineCategory)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, medicineCategory.Name, result.Name)
		assert.Equal(t, medicineCategory.Description, result.Description)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", medicineCategory).Return(nil, errors.New("Insert Process Failed")).Once()

		result, err := service.CreateMedicineCategory(medicineCategory)

		assert.Error(t, err)
		assert.EqualError(t, err, "Insert Process Failed")
		assert.Nil(t, result)
	})
}

func TestUpdateMedicineCategory(t *testing.T) {
	data := mocks.NewMedicineCategoryDataInterface(t)
	service := New(data)
	medicineCategory := medicinecategories.UpdateMedicineCategory{
		Name:        "Panadol",
		Description: "Panadol",
	}

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", medicineCategory, 1).Return(true, nil).Once()

		result, err := service.UpdateMedicineCategory(medicineCategory, 1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", medicineCategory, 1).Return(false, errors.New("Update Process Failed")).Once()

		result, err := service.UpdateMedicineCategory(medicineCategory, 1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Update Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestDeleteMedicineCategory(t *testing.T) {
	data := mocks.NewMedicineCategoryDataInterface(t)
	service := New(data)

	t.Run("Success Update", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		result, err := service.DeleteMedicineCategory(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed")).Once()

		result, err := service.DeleteMedicineCategory(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Delete Process Failed")
		assert.Equal(t, false, result)
	})
}
