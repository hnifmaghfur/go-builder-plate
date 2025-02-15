package models

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
