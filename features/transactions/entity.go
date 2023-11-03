package transactions

import "github.com/labstack/echo/v4"

type Transaction struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        int    `json:"amount"`
	SnapUrl       string `json:"snap_url"`
}

type TransactionInfo struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        int    `json:"amount"`
	SnapUrl       string `json:"snap_url"`
}

type TransactionInfoUser struct {
	IdentityNumber string `json:"identity_number"`
	UserName       string `json:"user_name"`
}

type TransactionDetail []struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        int    `json:"amount"`
	SnapUrl       string `json:"snap_url"`
}

type Payment struct {
	SnapUrl string `json:"snap_url"`
}

type TransactionDataInterface interface {
	Insert(idUser, idMcu int) (*Transaction, error)
	GetByID(idTrans int) (Transaction, error)
	UpdateSnapURL(snapUrl string, idTrans int) ([]Transaction, error)
	UpdateTransactionStatus(transID, status string) (bool, error)
	GetTransactions() ([]TransactionInfo, error)
	GetUserTransactions(idUser int) (TransactionInfoUser, TransactionDetail, error)
}

type TransactionServiceInterface interface {
	CreateTransaction(idUser, idMcu int) (*Transaction, error)
	PaymentTransaction(idTrans, idUser int) ([]Transaction, error)
	NotificationPayment(idTrans, idUser int) (string, error)
	GetTransactions() ([]TransactionInfo, error)
	GetUserTransactions(idUser int) (TransactionInfoUser, TransactionDetail, error)
}

type TransactionHandlerInterface interface {
	CreateTransaction() echo.HandlerFunc
	PaymentTransaction() echo.HandlerFunc
	NotificationPayment() echo.HandlerFunc
	GetTransactions() echo.HandlerFunc
}
