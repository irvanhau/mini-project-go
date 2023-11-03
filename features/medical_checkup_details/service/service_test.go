package service

import (
	medicalcheckupdetails "MiniProject/features/medical_checkup_details"
	"MiniProject/features/medical_checkup_details/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMedicalCheckupDetails(t *testing.T) {
	data := mocks.NewMedicalCheckupDetailDataInterface(t)
	service := New(data)
	medCheckupDetail := medicalcheckupdetails.MedicalCheckupDetailInfo{}
	medCheckupDetailInfo := medicalcheckupdetails.DetailInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", 1).Return(medCheckupDetail, medCheckupDetailInfo, nil).Once()

		resMed, resDetail, err := service.GetMedicalCheckupDetails(1)

		assert.Nil(t, err)
		assert.NotNil(t, resDetail)
		assert.NotNil(t, resMed)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", 1).Return(medCheckupDetail, medCheckupDetailInfo, errors.New("Get All Process Failed")).Once()

		resMed, resDetail, err := service.GetMedicalCheckupDetails(1)

		assert.Error(t, err)
		assert.Equal(t, medCheckupDetail, resMed)
		assert.Equal(t, medCheckupDetailInfo, resDetail)
		assert.EqualError(t, err, "Get All Process Failed")
	})
}

func TestGetMedicalCheckupDetail(t *testing.T) {
	data := mocks.NewMedicalCheckupDetailDataInterface(t)
	service := New(data)
	medCheckupDetail := []medicalcheckupdetails.MedicalCheckupDetailInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1, 1).Return(medCheckupDetail, nil).Once()

		resMed, err := service.GetMedicalCheckupDetail(1, 1)

		assert.Nil(t, err)
		assert.NotNil(t, resMed)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1, 1).Return(medCheckupDetail, errors.New("Get By ID Process Failed")).Once()

		resMed, err := service.GetMedicalCheckupDetail(1, 1)

		assert.Error(t, err)
		assert.Nil(t, resMed)
		assert.EqualError(t, err, "Get By ID Process Failed")
	})
}

func TestCreateMedicalCheckupDetail(t *testing.T) {
	data := mocks.NewMedicalCheckupDetailDataInterface(t)
	service := New(data)
	medCheckupDetail := medicalcheckupdetails.MedicalCheckupDetail{
		MedicalCheckupID: 1,
		MedicineID:       1,
		Quantity:         20,
	}

	t.Run("Success Insert", func(t *testing.T) {
		data.On("Insert", medCheckupDetail).Return(&medCheckupDetail, nil).Once()

		resMed, err := service.CreateMedicalCheckupDetail(medCheckupDetail)

		data.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, resMed)
		assert.Equal(t, medCheckupDetail.MedicalCheckupID, resMed.MedicalCheckupID)
		assert.Equal(t, medCheckupDetail.MedicineID, resMed.MedicineID)
		assert.Equal(t, medCheckupDetail.Quantity, resMed.Quantity)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", medCheckupDetail).Return(nil, errors.New("Insert Process Failed")).Once()

		resMed, err := service.CreateMedicalCheckupDetail(medCheckupDetail)

		assert.Error(t, err)
		assert.Nil(t, resMed)
		assert.EqualError(t, err, "Insert Process Failed")
	})
}

func TestUpdateMedicalCheckupDetail(t *testing.T) {
	data := mocks.NewMedicalCheckupDetailDataInterface(t)
	service := New(data)
	medCheckupDetail := medicalcheckupdetails.UpdateMedicalCheckupDetail{
		MedicineID: 1,
		Quantity:   20,
	}

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", medCheckupDetail, 1, 1).Return(true, nil).Once()

		resMed, err := service.UpdateMedicalCheckupDetail(medCheckupDetail, 1, 1)

		data.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, resMed)
		assert.Equal(t, true, resMed)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", medCheckupDetail, 1, 1).Return(false, errors.New("Update Process Failed")).Once()

		resMed, err := service.UpdateMedicalCheckupDetail(medCheckupDetail, 1, 1)

		assert.Error(t, err)
		assert.Equal(t, false, resMed)
		assert.EqualError(t, err, "Update Process Failed")
	})
}

func TestDeleteMedicalCheckupDetail(t *testing.T) {
	data := mocks.NewMedicalCheckupDetailDataInterface(t)
	service := New(data)

	t.Run("Success Delete", func(t *testing.T) {
		data.On("Delete", 1, 1).Return(true, nil).Once()

		resMed, err := service.DeleteMedicalCheckupDetail(1, 1)

		data.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, resMed)
		assert.Equal(t, true, resMed)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1, 1).Return(false, errors.New("Delete Process Failed")).Once()

		resMed, err := service.DeleteMedicalCheckupDetail(1, 1)

		assert.Error(t, err)
		assert.Equal(t, false, resMed)
		assert.EqualError(t, err, "Delete Process Failed")
	})
}
