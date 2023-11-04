package data

import (
	"MiniProject/features/transactions"
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) transactions.TransactionDataInterface {
	return &transactionData{
		db: db,
	}
}

// Insert implements transactions.TransactionDataInterface.
func (t *transactionData) Insert(idUser, idMcu int) (*transactions.Transaction, error) {
	var listMedicine []struct {
		MedicinePrice int `json:"medicine_price"`
		Quantity      int `json:"quantity"`
	}

	var TransactionIDInfo []struct {
		TransactionID string `json:"transaction_id"`
	}

	var total int
	TransactionID := uuid.NewString()

	var _ = t.db.Table("medical_checkup_details as mcd").
		Select("mcd.quantity as quantity", "m.price as medicine_price").
		Joins("JOIN medicines as m ON m.id = mcd.medicine_id").
		Where("mcd.medical_checkup_id = ?", idMcu).
		Where("mcd.deleted_at is null").
		Scan(&listMedicine)

	var _ = t.db.Table("transactions").
		Select("transaction_id").
		Where("user_id = ?", idUser).
		Where("deleted_at is null").
		Order("id DESC").
		First(&TransactionIDInfo)

	for _, medicine := range listMedicine {
		price := medicine.Quantity * medicine.MedicinePrice
		total += price
	}

	// for _, transID := range TransactionIDInfo {
	// 	TransactionNumber := transID.TransactionID[6:7]
	// 	number, _ := strconv.Atoi(TransactionNumber)
	// 	TransactionID = TransactionID
	// }

	var dbData = new(Transaction)
	dbData.TransactionID = TransactionID
	dbData.UserID = uint(idUser)
	dbData.Amount = total
	dbData.Status = "Pending"

	if err := t.db.Create(dbData).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return nil, err
	}

	var transaction = transactions.Transaction{
		ID:            dbData.ID,
		UserID:        dbData.UserID,
		Amount:        dbData.Amount,
		Status:        dbData.Status,
		TransactionID: dbData.TransactionID,
	}

	return &transaction, nil
}

func (t *transactionData) GetByID(idTrans int) (transactions.Transaction, error) {
	var listTransaction transactions.Transaction
	if err := t.db.Table("transactions").Where("id = ?", idTrans).Scan(&listTransaction).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return listTransaction, err
	}

	return listTransaction, nil
}

func (t *transactionData) UpdateSnapURL(snapUrl string, idTrans int) ([]transactions.Transaction, error) {
	var listTransaction = []transactions.Transaction{}
	if err := t.db.Table("transactions").Where("id = ?", idTrans).Updates(Transaction{
		SnapUrl: snapUrl,
	}).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return nil, err
	}

	if err := t.db.Table("transactions").Where("id = ?", idTrans).Scan(&listTransaction).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return nil, err
	}

	return listTransaction, nil

}

func (t *transactionData) UpdateTransactionStatus(transID, status string) (bool, error) {
	var qry = t.db.Table("transactions").Where("transaction_id = ?", transID).Updates(Transaction{Status: status})

	if err := qry.Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Failed, No Data Affected")
	}

	return true, nil
}

func (t *transactionData) GetTransactions() ([]transactions.TransactionInfo, error) {
	var listTransaction = []transactions.TransactionInfo{}
	var qry = t.db.Table("transactions").
		Select("users.full_name as user_name", "transactions.*").
		Joins("JOIN users ON users.id = transactions.user_id").
		Where("transactions.deleted_at is null")

	if err := qry.Scan(&listTransaction).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return nil, err
	}

	return listTransaction, nil
}

func (t *transactionData) GetUserTransactions(idUser int) (transactions.TransactionInfoUser, transactions.TransactionDetail, error) {
	var listTransaction transactions.TransactionInfoUser
	var listDetail transactions.TransactionDetail
	var qryTrans = t.db.Table("users").
		Select("users.full_name as user_name", "users.identity_number as identity_number").
		Joins("JOIN transactions ON users.id = transactions.user_id").
		Where("users.id = ?", idUser).
		Where("users.deleted_at is null")

	var qryDetail = t.db.Table("transactions").
		Select("transactions.*").
		Where("user_id = ?", idUser).
		Where("deleted_at is null").
		Scan(&listDetail)

	if err := qryTrans.Scan(&listTransaction).Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return listTransaction, listDetail, err
	}

	if err := qryDetail.Error; err != nil {
		log.Fatal("DB Error : ", err.Error())
		return listTransaction, listDetail, err
	}

	return listTransaction, listDetail, nil
}
