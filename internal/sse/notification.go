package sse

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/rabbitmq/amqp091-go"
)

type notificationStream struct {
	rmqConnection       *amqp091.Connection
	notificationService domain.NotificationService
	cnf                 *config.Config
}

func NewNotificationStream(rmqConnection *amqp091.Connection, notificationService domain.NotificationService, cnf *config.Config, app *fiber.App, atuhMidd fiber.Handler) {
	handler := notificationStream{
		rmqConnection:       rmqConnection,
		notificationService: notificationService,
		cnf:                 cnf,
	}

	app.Get("/sse/notification", atuhMidd, handler.StreamNotif)

}

func (n *notificationStream) StreamNotif(fctx *fiber.Ctx) error {

	//note buat function tersendiri untuk mekanisme get data from mQ

	fctx.Set("Content-Type", "text/event-stream")
	fctx.Set("Cache-Control", "no-cache")
	fctx.Set("Connection", "keep-alive")

	msgChan := make(chan dto.NotificationData)

	userId := fctx.Locals("x-user-id").(string)
	accountId := n.notificationService.FindAccountIdByUserId(fctx.Context(), userId)
	go n.notificationService.StreamNotif(fctx.Context(), accountId, msgChan)

	fctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		event := fmt.Sprintf("event: %s\n"+
			"data: \n\n", "initial")
		_, _ = fmt.Fprint(w, event)
		_ = w.Flush()
		for msg := range msgChan {
			if msg.AccountId == accountId {
				dataJson, _ := json.Marshal(msg)
				event = fmt.Sprintf("event: %s\n"+
					"data: %s\n\n", "notification-updated", dataJson)
				_, _ = fmt.Fprint(w, event)
				_ = w.Flush()
			}
		}
	})
	return nil
}
