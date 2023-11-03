package handler

type TransactionResponseDetail struct {
	UserName       string `json:"user_name"`
	IdentityNumber string `json:"identity_number"`
	DetailInfo     []struct {
		TransactionID string `json:"transaction_id"`
		Status        string `json:"status"`
		Amount        int    `json:"amount"`
	} `json:"transaction_detail"`
}
