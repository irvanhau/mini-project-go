package service

import (
	"MiniProject/features/appointments"
	"MiniProject/features/appointments/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppointments(t *testing.T) {
	data := mocks.NewAppointmentDataInterface(t)
	service := New(data)
	appointment := []appointments.AppointmentInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", 1).Return(appointment, 1, 1, nil).Once()

		res, totalData, totalPage, err := service.GetAppointments(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, totalData, 1)
		assert.Equal(t, totalPage, 1)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", 1).Return(nil, 0, 0, errors.New("Get All Process Failed"))

		res, totalData, totalPage, err := service.GetAppointments(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Failed")
		assert.Nil(t, res)
		assert.Equal(t, totalData, 0)
		assert.Equal(t, totalPage, 0)
	})
}

func TestGetAppointment(t *testing.T) {
	data := mocks.NewAppointmentDataInterface(t)
	service := New(data)
	appointment := []appointments.AppointmentInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(appointment, nil).Once()

		res, err := service.GetAppointment(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Failed"))

		res, err := service.GetAppointment(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get By ID Failed")
		assert.Nil(t, res)
	})
}

func TestCreateAppointment(t *testing.T) {
	data := mocks.NewAppointmentDataInterface(t)
	service := New(data)
	appointment := appointments.Appointment{
		UserID:          1,
		AppointmentDate: "2023-10-05",
		AppointmentTime: "17:00:00",
	}

	t.Run("Success Insert", func(t *testing.T) {
		data.On("Insert", appointment).Return(&appointment, nil).Once()

		res, err := service.CreateAppointment(appointment)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, appointment.UserID, res.UserID)
		assert.Equal(t, appointment.AppointmentDate, res.AppointmentDate)
		assert.Equal(t, appointment.AppointmentTime, res.AppointmentTime)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", appointment).Return(nil, errors.New("Insert Process Failed"))

		res, err := service.CreateAppointment(appointment)

		assert.Error(t, err)
		assert.EqualError(t, err, "Insert Process Failed")
		assert.Nil(t, res)
	})
}

func TestUpdateAppointment(t *testing.T) {
	data := mocks.NewAppointmentDataInterface(t)
	service := New(data)
	appointment := appointments.UpdateAppointment{
		AppointmentDate: "2023-10-05",
		AppointmentTime: "17:00:00",
	}

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", appointment, 1).Return(true, nil).Once()

		res, err := service.UpdateAppointment(appointment, 1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", appointment, 1).Return(false, errors.New("Update Process Failed"))

		res, err := service.UpdateAppointment(appointment, 1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Update Process Failed")
		assert.Equal(t, false, res)
	})
}

func TestDeleteAppointment(t *testing.T) {
	data := mocks.NewAppointmentDataInterface(t)
	service := New(data)

	t.Run("Success Delete", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		res, err := service.DeleteAppointment(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, true, res)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed"))

		res, err := service.DeleteAppointment(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Delete Process Failed")
		assert.Equal(t, false, res)
	})
}
