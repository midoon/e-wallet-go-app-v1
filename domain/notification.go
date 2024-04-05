package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"gorm.io/gorm"
)

type Notification struct {
	ID        string    `gorm:"column:id;primary_key"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	Status    int       `gorm:"column:status"`
	IsRead    int       `gorm:"column:is_read"`
	AccountId string    `gorm:"column:account_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (n *Notification) BeforeCreate(db *gorm.DB) error {
	if n.ID == "" {
		id := uuid.New().String()
		n.ID = id
	}

	return nil
}

type NotificationRepository interface {
	Insert(ctx context.Context, notification *Notification) error
	FindByUserAccount(ctx context.Context, accountId string) ([]Notification, error)
	Update(ctx context.Context, notification *Notification, notifId string) error
}

type NotificationService interface {
	FindByUserAccount(ctx context.Context, userId string) ([]dto.NotificationData, error)
	FindAccountIdByUserId(ctx context.Context, userId string) string
	StreamNotif(ctx context.Context, accountId string, msgChan chan<- dto.NotificationData)
}
