package model

type Wallet struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Status    bool   `json:"status"`
	Balance   int64  `json:"balance"`
	Point     int64  `json:"point"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatedBy string `json:"created_by"`
}
