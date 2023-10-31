package data

import (
	"gorm.io/gorm"
)

type Appointment struct {
	*gorm.Model
	UserID          uint   `gorm:"user_id"`
	AppointmentDate string `gorm:"column:appointment_date;type:date"`
	AppointmentTime string `gorm:"column:appointment_time;type:varchar(100)"`
}
