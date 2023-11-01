package data

import "gorm.io/gorm"

type MedicalCheckupDetail struct {
	*gorm.Model
	MedicalCheckupID uint `gorm:"medical_checkup_id"`
	MedicineID       uint `gorm:"medicine_id"`
	Quantity         int  `gorm:"column:quantity;type:int"`
}
