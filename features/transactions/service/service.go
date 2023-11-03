package service

import (
	"MiniProject/features/transactions"
	"MiniProject/utils/midtrans"
	"errors"
)

type transactionService struct {
	data transactions.TransactionDataInterface
	mid  midtrans.MidtransService
}

func New(d transactions.TransactionDataInterface, mid midtrans.MidtransService) transactions.TransactionServiceInterface {
	return &transactionService{
		data: d,
		mid:  mid,
	}
}

// ConfirmedTransaction implements transactions.TransactionServiceInterface.
// func (t *transactionService) ConfirmedTransaction(id string) error {
// 	panic("unimplemented")
// }

func (t *transactionService) CreateTransaction(idUser, idMcu int) (*transactions.Transaction, error) {
	result, err := t.data.Insert(idUser, idMcu)

	if err != nil {
		return nil, errors.New("Create Process Failed")
	}

	return result, nil
}

func (t *transactionService) PaymentTransaction(idTrans, idUser int) ([]transactions.Transaction, error) {
	resTrans, err := t.data.GetByID(idTrans)
	if err != nil {
		return nil, errors.New("Get By ID Transaction Failed")
	}

	snapURL, err := t.mid.GenerateSnapURL(resTrans)
	if err != nil {
		return nil, errors.New("Generate Snap URL Failed")
	}

	resUpdTrans, err := t.data.UpdateSnapURL(snapURL, idTrans)

	return resUpdTrans, nil
}

func (t *transactionService) NotificationPayment(idTrans, idUser int) (string, error) {
	resTrans, err := t.data.GetByID(idTrans)
	if err != nil {
		return "", errors.New("Get By ID Transaction Failed")
	}

	notif, err := t.mid.GetNotification(resTrans.TransactionID)
	if err != nil {
		return "", errors.New("Get Notification Payment Failed")
	}

	return notif, nil
}

func (t *transactionService) GetTransactions() ([]transactions.TransactionInfo, error) {
	res, err := t.data.GetTransactions()

	if err != nil {
		return nil, errors.New("Get Transactions Failed")
	}

	return res, nil
}

func (t *transactionService) GetUserTransactions(idUser int) (transactions.TransactionInfoUser, transactions.TransactionDetail, error) {
	resTrans, resDetail, err := t.data.GetUserTransactions(idUser)

	if err != nil {
		return resTrans, resDetail, errors.New("Get User Transactions Failed")
	}

	return resTrans, resDetail, nil
}
