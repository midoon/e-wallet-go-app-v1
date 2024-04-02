package repository

import (
	"context"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) domain.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

// FindByUser implements domain.NotificationRepository.
func (n *notificationRepository) FindByUserAccount(ctx context.Context, accountId string) ([]domain.Notification, error) {
	notifications := []domain.Notification{}
	err := n.db.WithContext(ctx).Where("account_id = ?", accountId).Find(&notifications).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return notifications, nil

}

// Insert implements domain.NotificationRepository.
func (n *notificationRepository) Insert(ctx context.Context, notification *domain.Notification) error {
	err := n.db.WithContext(ctx).Create(notification).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Update implements domain.NotificationRepository.
func (n *notificationRepository) Update(ctx context.Context, notification *domain.Notification, notifId string) error {
	err := n.db.WithContext(ctx).Where("id = ?", notifId).Updates(domain.Notification{
		IsRead: notification.IsRead,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
