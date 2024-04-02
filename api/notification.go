package api

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
)

type notificationApi struct {
	notificationService domain.NotificationService
}

func NewNotificationApi(notificationService domain.NotificationService, app *fiber.App, authMidd fiber.Handler) {
	handler := notificationApi{
		notificationService: notificationService,
	}

	app.Get("api/notification", authMidd, handler.GetNotification)

}

func (n *notificationApi) GetNotification(fctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(fctx.Context(), 15*time.Second)
	defer cancel()

	userId := fctx.Locals("x-user-id").(string)
	notifications, err := n.notificationService.FindByUser(c, userId)
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return fctx.Status(200).JSON(notifications)
}
