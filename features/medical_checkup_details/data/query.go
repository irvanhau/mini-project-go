package data

import (
	medicalcheckupdetails "MiniProject/features/medical_checkup_details"
	"errors"
	"log"

	"gorm.io/gorm"
)

type MedicalCheckupDetailData struct {
	db *gorm.DB
}

func New(db *gorm.DB) medicalcheckupdetails.MedicalCheckupDetailDataInterface {
	return &MedicalCheckupDetailData{
		db: db,
	}
}

func (mcdd *MedicalCheckupDetailData) GetAll(idMcu int) (medicalcheckupdetails.MedicalCheckupDetailInfo, medicalcheckupdetails.DetailInfo, error) {
	var listMCUD medicalcheckupdetails.MedicalCheckupDetailInfo
	var listMedicine medicalcheckupdetails.DetailInfo
	var qry = mcdd.db.Table("medical_checkup_details as mcd").
		Select("mcd.medical_checkup_id as medical_checkup_id", "mc.complain as complain", "mc.treatment as treatment", "mcd.quantity as quantity", "m.name as medicine_name").
		Joins("JOIN medical_checkups as mc ON mc.id = mcd.medical_checkup_id").
		Joins("JOIN medicines as m ON m.id = mcd.medicine_id").
		Where("mcd.medical_checkup_id = ?", idMcu).
		Where("mcd.deleted_at is null")

	var qryMedicine = mcdd.db.Table("medical_checkup_details as mcd").
		Select("mcd.quantity as quantity", "m.name as medicine_name").
		Joins("JOIN medicines as m ON m.id = mcd.medicine_id").
		Where("mcd.medical_checkup_id = ?", idMcu).
		Where("mcd.deleted_at is null").
		Scan(&listMedicine)

	if err := qry.Scan(&listMCUD).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return listMCUD, listMedicine, err
	}

	if err := qryMedicine.Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return listMCUD, listMedicine, err
	}

	return listMCUD, listMedicine, nil
}
func (mcdd *MedicalCheckupDetailData) GetByID(idMcu, idMcuDetail int) ([]medicalcheckupdetails.MedicalCheckupDetailInfo, error) {
	var listMCUD = []medicalcheckupdetails.MedicalCheckupDetailInfo{}
	var qry = mcdd.db.Table("medical_checkup_details as mcd").
		Select("mc.id as medical_checkup_id", "mc.complain as complain", "mc.treatment as treatment", "mcd.quantity as quantity", "m.name as medicine_name").
		Joins("JOIN medical_checkups as mc ON mc.id = mcd.medical_checkup_id").
		Joins("JOIN medicines as m ON m.id = mcd.medicine_id").
		Where("mcd.medical_checkup_id = ?", idMcu).
		Where("mcd.id = ?", idMcuDetail).
		Where("mcd.deleted_at is null")

	if err := qry.Scan(&listMCUD).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return nil, err
	}

	return listMCUD, nil
}
func (mcdd *MedicalCheckupDetailData) Insert(newData medicalcheckupdetails.MedicalCheckupDetail) (*medicalcheckupdetails.MedicalCheckupDetail, error) {
	var dbData = new(MedicalCheckupDetail)
	dbData.MedicalCheckupID = newData.MedicalCheckupID
	dbData.MedicineID = newData.MedicineID
	dbData.Quantity = newData.Quantity

	if err := mcdd.db.Create(dbData).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}
func (mcdd *MedicalCheckupDetailData) Update(newData medicalcheckupdetails.UpdateMedicalCheckupDetail, idMcu, idMcuDetail int) (bool, error) {
	var qry = mcdd.db.Table("medical_checkup_details").Where("id = ?", idMcuDetail).Where("medical_checkup_id = ?", idMcu).Updates(MedicalCheckupDetail{MedicineID: newData.MedicineID, Quantity: newData.Quantity})

	if err := qry.Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
func (mcdd *MedicalCheckupDetailData) Delete(idMcu, idMcuDetail int) (bool, error) {
	var deleteData = new(MedicalCheckupDetail)

	if err := mcdd.db.Where("id = ?", idMcuDetail).Where("medical_checkup_id  = ?", idMcu).Delete(&deleteData).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return false, err
	}

	return true, nil
}
