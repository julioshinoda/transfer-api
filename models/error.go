package models

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}
