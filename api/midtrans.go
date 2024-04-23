package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/domain"
)

type midtransApi struct {
	midtransService domain.MidtransService
	topupService    domain.TopupService
}

func NewMidtransApi(app *fiber.App, midtransService domain.MidtransService, topupService domain.TopupService) {
	m := midtransApi{
		midtransService: midtransService,
		topupService:    topupService,
	}

	app.Post("/midtrans/payment-callback", m.paymentHandlerNotification)
}

func (m *midtransApi) paymentHandlerNotification(fctx *fiber.Ctx) error {
	var notificationPayload map[string]interface{}

	if err := fctx.BodyParser(&notificationPayload); err != nil {
		return fctx.SendStatus(400)
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return fctx.SendStatus(400)
	}

	success, _ := m.midtransService.VerifyPayment(fctx.Context(), orderId)
	if success {
		_ = m.topupService.ConfirmedTopUp(fctx.Context(), orderId)
		return fctx.SendStatus(200)
	}

	return fctx.SendStatus(400)
}
