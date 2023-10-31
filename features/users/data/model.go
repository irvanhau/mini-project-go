package data

import (
	"MiniProject/features/appointments/data"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Email          string             `gorm:"column:email;unique;type:varchar(150)"`
	Password       string             `gorm:"column:password;unique;type:varchar(255)"`
	IdentityNumber string             `gorm:"column:identity_number;unique;type:varchar(255)"`
	FullName       string             `gorm:"column:full_name;type:varchar(150)"`
	BOD            string             `gorm:"column:bod;type:date"`
	Address        string             `gorm:"column:address;type:varchar(255)"`
	Role           string             `gorm:"column:role;type:enum('Admin','Patient')"`
	Appointments   []data.Appointment `gorm:"foreignKey:UserID"`
}
