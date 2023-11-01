package data

import (
	medicalcheckups "MiniProject/features/medical_checkups"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MedicalCheckupData struct {
	db *gorm.DB
}

func New(db *gorm.DB) medicalcheckups.MedicalCheckupDataInterface {
	return &MedicalCheckupData{
		db: db,
	}
}

func (mcd *MedicalCheckupData) GetAll() ([]medicalcheckups.MedicalCheckupInfo, error) {
	var listMCU = []medicalcheckups.MedicalCheckupInfo{}
	var qry = mcd.db.Table("medical_checkups").
		Select("medical_checkups.*", "users.full_name as user_name").
		Joins("JOIN users ON users.id = medical_checkups.user_id").
		Where("medical_checkups.deleted_at is null")

	if err := qry.Scan(&listMCU).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listMCU, nil

}
func (mcd *MedicalCheckupData) GetByID(id int) ([]medicalcheckups.MedicalCheckupInfo, error) {
	var listMCU = []medicalcheckups.MedicalCheckupInfo{}
	var qry = mcd.db.Table("medical_checkups").
		Select("medical_checkups.*", "users.full_name as user_name").
		Joins("JOIN users ON users.id = medical_checkups.user_id").
		Where("medical_checkups.deleted_at is null").
		Where("medical_checkups.id = ?", id)

	if err := qry.Scan(&listMCU).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listMCU, nil
}
func (mcd *MedicalCheckupData) Insert(newData medicalcheckups.MedicalCheckup) (*medicalcheckups.MedicalCheckup, error) {
	var dbData = new(MedicalCheckup)
	dbData.UserID = newData.UserID
	dbData.Complain = newData.Complain
	dbData.Treatment = newData.Treatment

	if err := mcd.db.Create(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}
func (mcd *MedicalCheckupData) Update(newData medicalcheckups.UpdateMedicalCheckup, id int) (bool, error) {
	var qry = mcd.db.Table("medical_checkups").Where("id = ?", id).Updates(MedicalCheckup{
		Complain:  newData.Complain,
		Treatment: newData.Treatment,
	})

	if err := qry.Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
func (mcd *MedicalCheckupData) Delete(id int) (bool, error) {
	var deleteData = new(MedicalCheckup)

	if err := mcd.db.Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return false, err
	}

	return true, nil
}
