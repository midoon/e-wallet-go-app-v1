package dto

import "time"

type NotificationData struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Status    int       `json:"status"`
	IsRead    int       `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
