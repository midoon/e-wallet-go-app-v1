package dto

type NotificationResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []NotificationData `json:"data"`
}
