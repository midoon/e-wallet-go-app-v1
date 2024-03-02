package dto

type TokenData struct {
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    TokenData `json:"data"`
}
