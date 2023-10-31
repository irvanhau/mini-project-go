package handler

type InputResponse struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}
