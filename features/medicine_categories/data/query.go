package data

import (
	medicinecategories "MiniProject/features/medicine_categories"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MedicineCategoryData struct {
	db *gorm.DB
}

func New(db *gorm.DB) medicinecategories.MedicineCategoryDataInterface {
	return &MedicineCategoryData{
		db: db,
	}
}

func (mcd *MedicineCategoryData) GetAll() ([]medicinecategories.MedicineCategory, error) {
	var listMedicineCategory = []medicinecategories.MedicineCategory{}
	if err := mcd.db.Find(&listMedicineCategory).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listMedicineCategory, nil
}
func (mcd *MedicineCategoryData) GetByID(id int) ([]medicinecategories.MedicineCategory, error) {
	var listMedicineCategory = []medicinecategories.MedicineCategory{}
	if err := mcd.db.Where("id = ?", id).First(&listMedicineCategory).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listMedicineCategory, nil
}
func (mcd *MedicineCategoryData) Insert(newData medicinecategories.MedicineCategory) (*medicinecategories.MedicineCategory, error) {
	var dbData = new(MedicineCategory)
	dbData.Name = newData.Name
	dbData.Description = newData.Description

	if err := mcd.db.Create(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}
func (mcd *MedicineCategoryData) Update(newData medicinecategories.UpdateMedicineCategory, id int) (bool, error) {
	var qry = mcd.db.Table("medicine_categories").Where("id = ?", id).Updates(MedicineCategory{Name: newData.Name, Description: newData.Description})

	if err := qry.Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
func (mcd *MedicineCategoryData) Delete(id int) (bool, error) {
	var deletedata = new(MedicineCategory)

	if err := mcd.db.Where("id = ?", id).Delete(&deletedata).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return false, err
	}

	return true, nil
}
