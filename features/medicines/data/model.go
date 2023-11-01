package data

import (
	"MiniProject/features/medical_checkup_details/data"

	"gorm.io/gorm"
)

type Medicine struct {
	*gorm.Model
	CategoryID           uint                        `gorm:"category_id"`
	Name                 string                      `gorm:"column:name;type:varchar(255)"`
	StockMinimum         int                         `gorm:"column:stock_minimum;type:int"`
	Stock                int                         `gorm:"column:stock;type:int"`
	Price                int                         `gorm:"column:price;type:int"`
	Photo                string                      `gorm:"column:photo;type:varchar(255)"`
	File                 string                      `gorm:"column:file;type:varchar(255)"`
	MedicalCheckupDetail []data.MedicalCheckupDetail `gorm:"foreignKey:MedicineID"`
}
