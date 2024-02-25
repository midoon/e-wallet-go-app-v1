package dto

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
