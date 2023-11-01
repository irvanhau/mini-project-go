package handler

type InputResponse struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	Complain  string `json:"complain" form:"complain"`
	Treatment string `json:"treatment" form:"treatment"`
}

type UpdateResponse struct {
	Complain  string `json:"complain" form:"complain"`
	Treatment string `json:"treatment" form:"treatment"`
}
