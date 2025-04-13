package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type RegisterRequest struct {
	Email          string `json:"email" validate:"required,email"`
	PhoneNumber    string `json:"phone_number" validate:"required,max=15"`
	Password       string `json:"password" validate:"required"`
	RetypePassword string `json:"retype_password" validate:"required"`
}

type MissionRequest struct {
	Name      string      `json:"name" validate:"required"`
	Type      string      `json:"type" validate:"required"`
	Point     int         `json:"point" validate:"required"`
	Condition interface{} `json:"condition,omitempty" validate:"required"`
}
