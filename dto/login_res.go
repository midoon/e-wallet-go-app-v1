package dto

type LoginResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    TokenData `json:"data"`
}
