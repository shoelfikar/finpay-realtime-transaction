package model

type User struct {
	Id           int     `json:"id"`
	Password     *string `json:"password,omitempty"`
	Email        string  `json:"email"`
	Status       int     `json:"status"`
	Role         string  `json:"role"`
	Token        *string `json:"token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	CreatedBy    string  `json:"created_by"`
}
