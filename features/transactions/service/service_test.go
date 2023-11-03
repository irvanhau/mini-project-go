package service

import (
	"MiniProject/features/transactions"
	"MiniProject/features/transactions/mocks"
	midtrans "MiniProject/utils/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	mid := midtrans.NewMidtransService(t)
	service := New(data, mid)
	transaction := transactions.Transaction{
		UserID:        1,
		TransactionID: "Trans-1",
		Status:        "Pending",
		Amount:        20000,
		SnapUrl:       "http.midtrans.com",
	}

	t.Run("Success Create", func(t *testing.T) {
		data.On("Insert", 1, 1).Return(&transaction, nil).Once()

		res, err := service.CreateTransaction(1, 1)

		data.AssertExpectations(t)
		mid.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, transaction.UserID, res.UserID)
		assert.Equal(t, transaction.TransactionID, res.TransactionID)
		assert.Equal(t, transaction.Status, res.Status)
		assert.Equal(t, transaction.Amount, res.Amount)
		assert.Equal(t, transaction.SnapUrl, res.SnapUrl)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", 1, 1).Return(nil, errors.New("Create Process Failed")).Once()

		res, err := service.CreateTransaction(1, 1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Create Process Failed")
	})
}

func TestPaymentTransaction(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	mid := midtrans.NewMidtransService(t)
	service := New(data, mid)
	transactionStruct := []transactions.Transaction{}
	var transaction transactions.Transaction

	t.Run("Success Payment", func(t *testing.T) {
		data.On("GetByID", 1).Return(transaction, nil).Once()
		mid.On("GenerateSnapURL", transaction).Return("http.midtrans.com", nil).Once()
		data.On("UpdateSnapURL", "http.midtrans.com", 1).Return(transactionStruct, nil).Once()

		res, err := service.PaymentTransaction(1, 1)

		data.AssertExpectations(t)
		mid.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Get By ID Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(transaction, errors.New("Get By ID Transaction Failed")).Once()

		res, err := service.PaymentTransaction(1, 1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Get By ID Transaction Failed")
	})

	t.Run("Generate Snap URL Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(transaction, nil).Once()
		mid.On("GenerateSnapURL", transaction).Return("", errors.New("Generate Snap URL Failed")).Once()

		res, err := service.PaymentTransaction(1, 1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Generate Snap URL Failed")
	})

	t.Run("Update Snap URL Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(transaction, nil).Once()
		mid.On("GenerateSnapURL", transaction).Return("http.midtrans.com", nil).Once()
		data.On("UpdateSnapURL", "http.midtrans.com", 1).Return(nil, errors.New("Update Snap URL Failed")).Once()

		res, err := service.PaymentTransaction(1, 1)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "Update Snap URL Failed")
	})
}

func TestNotificationPayment(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	mid := midtrans.NewMidtransService(t)
	service := New(data, mid)
	var transaction transactions.Transaction

	t.Run("Success Notification", func(t *testing.T) {
		data.On("GetByID", 1).Return(transaction, nil).Once()
		mid.On("GetNotification", transaction.TransactionID).Return("Success", nil).Once()

		result, err := service.NotificationPayment(1, 1)

		data.AssertExpectations(t)
		mid.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Get By ID Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(transaction, errors.New("Get By ID Transaction Failed")).Once()

		res, err := service.NotificationPayment(1, 1)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.EqualError(t, err, "Get By ID Transaction Failed")
	})

	t.Run("Notification Failed", func(t *testing.T) {
		data.On("GetByID", 1).Return(transaction, nil).Once()
		mid.On("GetNotification", transaction.TransactionID).Return("", errors.New("Get Notification Payment Failed")).Once()

		res, err := service.NotificationPayment(1, 1)

		assert.Error(t, err)
		assert.Equal(t, "", res)
		assert.EqualError(t, err, "Get Notification Payment Failed")
	})
}

func TestGetTransactions(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	mid := midtrans.NewMidtransService(t)
	service := New(data, mid)
	var transaction = []transactions.TransactionInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetTransactions").Return(transaction, nil).Once()

		result, err := service.GetTransactions()

		data.AssertExpectations(t)
		mid.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetTransactions").Return(nil, errors.New("Get Transactions Failed")).Once()

		result, err := service.GetTransactions()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "Get Transactions Failed")
	})
}

func TestGetUserTransactions(t *testing.T) {
	data := mocks.NewTransactionDataInterface(t)
	mid := midtrans.NewMidtransService(t)
	service := New(data, mid)
	var transUser transactions.TransactionInfoUser
	var transDetail transactions.TransactionDetail

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetUserTransactions", 1).Return(transUser, transDetail, nil).Once()

		resUser, resDetail, err := service.GetUserTransactions(1)

		data.AssertExpectations(t)
		mid.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, transDetail, resDetail)
		assert.Equal(t, transUser, resUser)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetUserTransactions", 1).Return(transUser, transDetail, errors.New("Get User Transactions Failed")).Once()

		resUser, resDetail, err := service.GetUserTransactions(1)

		assert.Error(t, err)
		assert.Equal(t, transDetail, resDetail)
		assert.Equal(t, transUser, resUser)
		assert.EqualError(t, err, "Get User Transactions Failed")
	})

}
