package service

import (
	"context"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
}

// FindByUser implements domain.NotificationService.

func NewNotificationService(notificationRepository domain.NotificationRepository) domain.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
	}
}

func (n *notificationService) FindByUser(ctx context.Context, userId string) ([]dto.NotificationData, error) {
	notifications := []dto.NotificationData{}
	notifs, err := n.notificationRepository.FindByUser(ctx, userId)
	if err != nil {
		return []dto.NotificationData{}, err
	}
	for _, val := range notifs {
		notifications = append(notifications, dto.NotificationData{
			ID:        val.ID,
			Title:     val.Title,
			Body:      val.Body,
			Status:    val.Status,
			IsRead:    val.IsRead,
			CreatedAt: val.CreatedAt,
		})
	}

	return notifications, nil
}
