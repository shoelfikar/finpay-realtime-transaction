package model

type User struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	Status    int     `json:"status"`
	Role      string  `json:"role"`
	Token     *string `json:"token,omitempty"`
	CreatedAt string  `json:"created_at"`
	CreatedBy string  `json:"created_by"`
}
