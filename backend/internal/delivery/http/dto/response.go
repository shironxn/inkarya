package dto

type Response struct {
	Success bool         `json:"success"`
	Status  int          `json:"status"`
	Message string       `json:"message,omitempty"`
	Errors  []FieldError `json:"errors,omitempty"`
	Data    interface{}  `json:"data,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
