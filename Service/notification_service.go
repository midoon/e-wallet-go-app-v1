package service

import (
	"context"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
	accountRepository      domain.AccountRepository
}

// FindByUser implements domain.NotificationService.

func NewNotificationService(notificationRepository domain.NotificationRepository, accountRepository domain.AccountRepository) domain.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
		accountRepository:      accountRepository,
	}
}

func (n *notificationService) FindByUserAccount(ctx context.Context, userId string) ([]dto.NotificationData, error) {

	// get accound by userId
	account, err := n.accountRepository.FindByUserId(ctx, userId)
	if err != nil {
		return []dto.NotificationData{}, err
	}

	notifications := []dto.NotificationData{}
	notifs, err := n.notificationRepository.FindByUserAccount(ctx, account.ID)
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
