package database

import (
	dataAppointment "MiniProject/features/appointments/data"
	dataMCUDetail "MiniProject/features/medical_checkup_details/data"
	dataMCU "MiniProject/features/medical_checkups/data"
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
	db.AutoMigrate(dataMCU.MedicalCheckup{})
	db.AutoMigrate(dataMCUDetail.MedicalCheckupDetail{})
}
