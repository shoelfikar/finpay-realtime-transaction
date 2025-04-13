package model

type User struct {
	Id           string  `json:"id"`
	Password     *string `json:"password,omitempty"`
	Email        string  `json:"email"`
	PhoneNumber  string  `json:"phone_number"`
	Status       int     `json:"status"`
	Role         string  `json:"role"`
	Balance      int     `json:"balance"`
	Point        int     `json:"point"`
	Token        *string `json:"token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	CreatedBy    string  `json:"created_by"`
}
