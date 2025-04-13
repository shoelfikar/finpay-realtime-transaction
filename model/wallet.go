package model

type Wallet struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Status    bool   `json:"status"`
	Balance   int64  `json:"balance"`
	Point     int64  `json:"point"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatedBy string `json:"created_by"`
}
