package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID            string    `gorm:"primary_key;column:id"`
	AccountNumber string    `gorm:"uniqueIndex;column:account_number"`
	Pin           string    `gorm:"column:pin"`
	Balance       float64   `gorm:"column:balance"`
	UserId        string    `gorm:"uniqueIndex;column:user_id"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime"`
}

func (a *Account) BeforeCreate(db *gorm.DB) error {
	if a.ID == "" {
		id := uuid.New().String()
		a.ID = id
	}

	return nil
}

type AccountRepository interface {
	Insert(ctx context.Context, account *Account) error
}
