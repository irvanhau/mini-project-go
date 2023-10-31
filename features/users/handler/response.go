package handler

type RegisterResponse struct {
	Email          string `json:"email" form:"email"`
	IdentityNumber string `json:"identity_number" form:"identity_number"`
	FullName       string `json:"full_name" form:"full_name"`
	BOD            string `json:"bod" form:"bod"`
	Address        string `json:"address" form:"address"`
	Role           string `json:"role" form:"role"`
}

type LoginResponse struct {
	Email          string `json:"email" form:"email"`
	IdentityNumber string `json:"identity_number" form:"identity_number"`
	FullName       string `json:"full_name" form:"full_name"`
	BOD            string `json:"bod" form:"bod"`
	Address        string `json:"address" form:"address"`
	Token          any    `json:"token"`
}
