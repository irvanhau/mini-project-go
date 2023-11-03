package midtrans

import (
	"MiniProject/configs"
	"MiniProject/features/transactions"
	"errors"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	GenerateSnapURL(transactions.Transaction) (string, error)
	GetNotification(transID string) (string, error)
}

type midtarnsService struct {
	client  snap.Client
	configs configs.ProgramConfig
	data    transactions.TransactionDataInterface
}

func NewMidtrans(cfg configs.ProgramConfig, data transactions.TransactionDataInterface) MidtransService {
	var client snap.Client
	envi := midtrans.Sandbox
	client.New(cfg.MTServerKey, envi)

	return &midtarnsService{
		client:  client,
		configs: cfg,
		data:    data,
	}
}

// GenerateSnapURL implements MidtransService.
func (m *midtarnsService) GenerateSnapURL(t transactions.Transaction) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  t.TransactionID,
			GrossAmt: int64(t.Amount),
		},
	}

	snapResp, err := m.client.CreateTransaction(req)
	if err != nil {
		return "", err
	}

	snapURL := snapResp.RedirectURL

	return snapURL, nil
}

func (m *midtarnsService) GetNotification(transID string) (string, error) {
	var c coreapi.Client
	c.New(m.client.ServerKey, midtrans.Sandbox)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := c.CheckTransaction(transID)
	if e != nil {
		return "", errors.New("Trans ID not found")
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
				res, err := m.data.UpdateTransactionStatus(transID, "Success")
				if err != nil {
					return "", errors.New("Update Transaction Status Failed")
				}

				if res != true {
					return "", nil
				}

				return "Transaction Success", nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
				res, err := m.data.UpdateTransactionStatus(transID, "Cancel")
				if err != nil {
					return "", errors.New("Update Transaction Status Failed")
				}

				if res != true {
					return "", nil
				}

				return "Transaction Expired", nil
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
				res, err := m.data.UpdateTransactionStatus(transID, "Pending")
				if err != nil {
					return "", errors.New("Update Transaction Status Failed")
				}

				if res != true {
					return "", nil
				}

				return "Transaction Pending", nil
			}
		}
	}
	return "", nil
}
