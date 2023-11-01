package data

import (
	"MiniProject/features/medical_checkup_details/data"

	"gorm.io/gorm"
)

type MedicalCheckup struct {
	*gorm.Model
	UserID               uint                        `gorm:"user_id"`
	Complain             string                      `gorm:"column:complain;type:varchar(255)"`
	Treatment            string                      `gorm:"column:treatment;type:varchar(255)"`
	MedicalCheckupDetail []data.MedicalCheckupDetail `gorm:"foreignKey:MedicalCheckupID"`
}
