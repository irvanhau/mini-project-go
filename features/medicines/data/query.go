package data

import (
	"MiniProject/features/medicines"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MedicineData struct {
	db *gorm.DB
}

func New(db *gorm.DB) medicines.MedicineDataInterface {
	return &MedicineData{
		db: db,
	}
}

func (md *MedicineData) GetAll(kategori int, name string) ([]medicines.MedicineInfo, error) {
	var listMedicine = []medicines.MedicineInfo{}

	var qry = md.db.Table("medicines").Select("medicines.*", "medicine_categories.name as category_name", "medicine_categories.id as category_id").
		Joins("JOIN medicine_categories ON medicines.category_id = medicine_categories.id").
		Where("medicines.deleted_at is null")

	if kategori != 0 {
		qry.Where("category_id = ?", kategori)
	}

	if name != "" {
		qry.Where("medicines.name LIKE ?", "%"+name+"%")
	}

	if err := qry.Scan(&listMedicine).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listMedicine, nil
}
func (md *MedicineData) GetByID(id int) ([]medicines.MedicineInfo, error) {
	var listMedicine = []medicines.MedicineInfo{}
	var qry = md.db.Table("medicines").Select("medicines.*", "medicine_categories.name as category_name").
		Joins("JOIN medicine_categories ON medicines.category_id = medicine_categories.id").
		Where("medicines.deleted_at is null").
		Where("medicines.id = ?", id).
		Scan(&listMedicine)

	if err := qry.Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listMedicine, nil

}
func (md *MedicineData) Insert(newData medicines.Medicine) (*medicines.Medicine, error) {
	var dbData = new(Medicine)
	dbData.CategoryID = newData.CategoryID
	dbData.Name = newData.Name
	dbData.StockMinimum = newData.StockMinimum
	dbData.Stock = newData.Stock
	dbData.Price = newData.Price
	dbData.Photo = newData.Photo
	dbData.File = newData.File

	if err := md.db.Create(&newData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}
func (md *MedicineData) Update(newData medicines.UpdateMedicine, id int) (bool, error) {
	var qry = md.db.Table("medicines").
		Where("id = ?", id).
		Updates(Medicine{
			Name:         newData.Name,
			CategoryID:   newData.CategoryID,
			StockMinimum: newData.StockMinimum,
			Stock:        newData.Stock,
			Price:        newData.Price,
		})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
func (md *MedicineData) Delete(id int) (bool, error) {
	var deleteData = new(Medicine)

	if err := md.db.Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (md *MedicineData) UpdateFileMedicine(file string, id int) (bool, error) {
	var qry = md.db.Table("medicines").
		Where("id = ?", id).
		Updates(Medicine{File: file})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (md *MedicineData) UpdatePhotoMedicine(photo string, id int) (bool, error) {
	var qry = md.db.Table("medicines").
		Where("id = ?", id).
		Updates(Medicine{Photo: photo})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
