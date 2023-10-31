package data

import (
	"MiniProject/features/appointments"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppointmentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) appointments.AppointmentDataInterface {
	return &AppointmentData{
		db: db,
	}
}

func (ad *AppointmentData) GetAll(page int) ([]appointments.AppointmentInfo, int, int, error) {
	var listAppointment = []appointments.AppointmentInfo{}

	var limit = 3
	if page == 0 {
		limit = -1
	}

	var qry = ad.db.Table("appointments").
		Select("appointments.*", "users.full_name as user_name").
		Joins("JOIN users ON users.id = appointments.user_id").
		Where("appointments.deleted_at is null").
		Where("appointments.appointment_date >= CURDATE()").
		Where("appointments.appointment_time >= CURTIME()").
		Offset((page - 1) * 3).
		Order("appointments.appointment_date ASC").
		Order("appointments.appointment_time ASC").
		Limit(limit).
		Scan(&listAppointment)

	totalData, totalPage := ad.Pagination()

	if err := qry.Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, 0, 0, err
	}

	return listAppointment, totalData, totalPage, nil
}
func (ad *AppointmentData) GetByID(id int) ([]appointments.AppointmentInfo, error) {
	var listAppointment = []appointments.AppointmentInfo{}
	var qry = ad.db.Table("appointments").
		Select("appointments.*", "users.full_name as user_name").
		Joins("JOIN users ON users.id = appointments.user_id").
		Where("appointments.deleted_at is null").
		Where("appointments.id = ?", id).
		Scan(&listAppointment)

	if err := qry.Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listAppointment, nil
}
func (ad *AppointmentData) Insert(newData appointments.Appointment) (*appointments.Appointment, error) {
	var dbData = new(Appointment)
	dbData.UserID = newData.UserID
	dbData.AppointmentDate = newData.AppointmentDate
	dbData.AppointmentTime = newData.AppointmentTime

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ad *AppointmentData) Update(newData appointments.UpdateAppointment, id int) (bool, error) {
	var qry = ad.db.Table("appointments").Where("id = ?", id).Updates(Appointment{
		AppointmentDate: newData.AppointmentDate,
		AppointmentTime: newData.AppointmentTime,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (ad *AppointmentData) Delete(id int) (bool, error) {
	var deleteData = new(Appointment)

	if err := ad.db.Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (ad *AppointmentData) Pagination() (int, int) {
	var listAppointment = []appointments.AppointmentInfo{}
	var count int64

	var _ = ad.db.Table("appointments").
		Select("appointments.*", "users.full_name as user_name").
		Joins("JOIN users ON users.id = appointments.user_id").
		Where("appointments.deleted_at is null").
		Where("appointments.appointment_date >= CURDATE()").
		Where("appointments.appointment_time >= CURTIME()").
		Order("appointments.appointment_date ASC").
		Order("appointments.appointment_time ASC").
		Count(&count).
		Scan(&listAppointment)

	countInt := int(count)

	totalPage := countInt / 3

	return countInt, totalPage
}
