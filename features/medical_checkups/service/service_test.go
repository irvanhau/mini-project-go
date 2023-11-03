package service

import (
	medicalcheckups "MiniProject/features/medical_checkups"
	"MiniProject/features/medical_checkups/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMedicalCheckups(t *testing.T) {
	data := mocks.NewMedicalCheckupDataInterface(t)
	service := New(data)
	medCheckup := []medicalcheckups.MedicalCheckupInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(medCheckup, nil).Once()

		result, err := service.GetMedicalCheckups()
		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		result, err := service.GetMedicalCheckups()
		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Process Failed")
		assert.Nil(t, result)
	})
}

func TestGetMedicalCheckup(t *testing.T) {
	data := mocks.NewMedicalCheckupDataInterface(t)
	service := New(data)
	medCheckup := []medicalcheckups.MedicalCheckupInfo{}

	t.Run("Success Get By ID", func(t *testing.T) {
		data.On("GetByID", 0).Return(medCheckup, nil).Once()

		result, err := service.GetMedicalCheckup(0)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 0).Return(nil, errors.New("Get By ID Process Failed")).Once()

		result, err := service.GetMedicalCheckup(0)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get By ID Process Failed")
		assert.Nil(t, result)
	})
}

func TestCreateMedicalCheckup(t *testing.T) {
	data := mocks.NewMedicalCheckupDataInterface(t)
	service := New(data)
	medCheckup := medicalcheckups.MedicalCheckup{
		UserID:    1,
		Complain:  "Mual",
		Treatment: "Istirahat",
	}

	t.Run("Success Insert", func(t *testing.T) {
		data.On("Insert", medCheckup).Return(&medCheckup, nil).Once()

		result, err := service.CreateMedicalCheckup(medCheckup)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, medCheckup.UserID, result.UserID)
		assert.Equal(t, medCheckup.Complain, result.Complain)
		assert.Equal(t, medCheckup.Treatment, result.Treatment)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", medCheckup).Return(nil, errors.New("Insert Process Failed"))

		result, err := service.CreateMedicalCheckup(medCheckup)

		assert.Error(t, err)
		assert.EqualError(t, err, "Insert Process Failed")
		assert.Nil(t, result)
	})
}

func TestUpdateMedicalCheckup(t *testing.T) {
	data := mocks.NewMedicalCheckupDataInterface(t)
	service := New(data)
	medCheckup := medicalcheckups.UpdateMedicalCheckup{
		Complain:  "Mual",
		Treatment: "Istirahat",
	}

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", medCheckup, 1).Return(true, nil).Once()

		result, err := service.UpdateMedicalCheckup(medCheckup, 1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", medCheckup, 1).Return(false, errors.New("Update Process Failed"))

		result, err := service.UpdateMedicalCheckup(medCheckup, 1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Update Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestDeleteMedicalCheckup(t *testing.T) {
	data := mocks.NewMedicalCheckupDataInterface(t)
	service := New(data)

	t.Run("Success Delete", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		result, err := service.DeleteMedicalCheckup(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed"))

		result, err := service.DeleteMedicalCheckup(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Delete Process Failed")
		assert.Equal(t, false, result)
	})
}
