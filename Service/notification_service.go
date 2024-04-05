package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/rabbitmq/amqp091-go"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
	accountRepository      domain.AccountRepository
	mqConnection           *amqp091.Connection
	cnf                    *config.Config
}

// FindByUser implements domain.NotificationService.

func NewNotificationService(notificationRepository domain.NotificationRepository, accountRepository domain.AccountRepository, mq *amqp091.Connection, cnf *config.Config) domain.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
		accountRepository:      accountRepository,
		mqConnection:           mq,
		cnf:                    cnf,
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
			AccountId: val.AccountId,
			CreatedAt: val.CreatedAt,
		})
	}

	return notifications, nil
}

func (n *notificationService) FindAccountIdByUserId(ctx context.Context, userId string) string {
	account, _ := n.accountRepository.FindByUserId(ctx, userId)
	return account.ID
}

func (n *notificationService) StreamNotif(ctx context.Context, accountId string, msgChan chan<- dto.NotificationData) {
	mqChannel, err := n.mqConnection.Channel()
	if err != nil {
		log.Println(err)
	}

	defer mqChannel.Close()

	messages, err := mqChannel.ConsumeWithContext(context.Background(), n.cnf.RabbitMQ.Queue, accountId, false, false, false, false, nil)
	if err != nil {
		log.Println(err)
	}

	for msg := range messages {
		var dataNotif dto.NotificationData
		_ = json.Unmarshal(msg.Body, &dataNotif)
		if dataNotif.AccountId == accountId {
			msg.Ack(false)
			msgChan <- dataNotif
		}
	}
}
