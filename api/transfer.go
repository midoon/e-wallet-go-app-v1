package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/domain"
)

type transferApi struct {
	transactionService domain.TransactionService
}

func NewTranferApi(app *fiber.App, transactionService domain.TransactionService, authMidd fiber.Handler) {
	handler := transferApi{
		transactionService: transactionService,
	}

	app.Post("/api/transfer/inquiry", authMidd, handler.transferInquiry)
	app.Post("/api/transfer/execute", authMidd, handler.transferExecute)
}

func (t *transferApi) transferInquiry(fctx *fiber.Ctx) error {
	return nil
}

func (t *transferApi) transferExecute(fctx *fiber.Ctx) error {
	return nil
}
