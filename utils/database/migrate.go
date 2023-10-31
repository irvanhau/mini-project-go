package database

import (
	dataAppointment "MiniProject/features/appointments/data"
	dataMedicineCategory "MiniProject/features/medicine_categories/data"
	dataMedicine "MiniProject/features/medicines/data"
	dataUser "MiniProject/features/users/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(dataUser.User{})
	db.AutoMigrate(dataMedicineCategory.MedicineCategory{})
	db.AutoMigrate(dataMedicine.Medicine{})
	db.AutoMigrate(dataAppointment.Appointment{})
}
