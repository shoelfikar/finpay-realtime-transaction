package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type RegisterRequest struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
}
