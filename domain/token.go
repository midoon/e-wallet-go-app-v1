package domain

import "time"

type Token struct {
	ID           string    `gorm:"primary_key;column:id"`
	RefreshToken string    `gorm:"column:refresh_token"`
	UserId       string    `gorm:"column:user_id;uniqueIndex"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime"`
}
