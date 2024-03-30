package dto

import "time"

type NotificationData struct {
	ID        string    `gorm:"column:id;primary_key"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	Status    int       `gorm:"column:status"`
	IsRead    int       `gorm:"column:is_read"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}
