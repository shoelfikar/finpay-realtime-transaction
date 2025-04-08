package model

type ValidationError struct {
	Errors []*ValidationMessage `json:"errors,omitempty"`
}

type ValidationMessage struct {
	Field string      `json:"field,omitempty"`
	Tag   string      `json:"tag,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Error string      `json:"error,omitempty"`
}
