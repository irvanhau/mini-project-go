package data

import (
	"MiniProject/features/medicines/data"

	"gorm.io/gorm"
)

type MedicineCategory struct {
	*gorm.Model
	Name               string          `gorm:"column:name;type:varchar(255)"`
	Description        string          `gorm:"column:description;unique;type:varchar(150)"`
	MedicineCategories []data.Medicine `gorm:"foreignKey:CategoryID"`
}
