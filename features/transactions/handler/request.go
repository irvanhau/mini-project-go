package handler

type InputRequest struct {
	ID      string `json:"id"`
	UserID  uint   `json:"user_id"`
	Status  string `json:"status"`
	Amount  int    `json:"amount"`
	SnapUrl string `json:"snap_url"`
}
