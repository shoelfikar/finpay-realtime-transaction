package model

type Transaction struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Balance    int64  `json:"balance"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Created_by string `json:"created_by"`
}

type TransactionDetail struct {
	Id            int    `json:"id"`
	TransactionId int    `json:"transaction_id"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	Description   string `json:"description"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Created_by    string `json:"created_by"`
}
