package model

type ResponseJSON struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success string      `json:"success"`
	Error   string      `json:"error,omitempty"`
}
