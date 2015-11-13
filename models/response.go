package models

// Response struct
type Response struct {
	Data      interface{} `json:"data,omitempty"`
	ErrorCode int         `json:"error_code,omitempty"`
	Success   bool        `json:"success"`
}
