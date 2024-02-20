package domain

import "time"

type User struct {
	ID        string    `gorm:"primary_key;column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email;uniqueIndex"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime"`
	Token     Token     `gorm:"foreignKey:user_id;referenceid"`
}
