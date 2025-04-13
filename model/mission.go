package model

type Missions struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Point     int         `json:"reward"`
	Condition interface{} `json:"condition"`
	Status    bool        `json:"status"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	CreatedBy string      `json:"created_by"`
}

type MissionCriteria struct {
	Id        int    `json:"id"`
	MissionId int    `json:"mission_id"`
	Type      string `json:"type"`
	Criteria  string `json:"criteria"`
}

type LoyaltyPoints struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	SourceType  string `json:"source_type"`
	SourceId    string `json:"source_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CreatedBy   string `json:"created_by"`
}
