package model

type Mission struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	RewardType string `json:"reward_type"`
	Reward     int    `json:"reward"`
	Status     int8   `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	CreatedBy  string `json:"created_by"`
}
