package dto

type RefreshData struct {
	AccessToken string `json:"access_token"`
}

type RefreshResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    RefreshData `json:"data"`
}
