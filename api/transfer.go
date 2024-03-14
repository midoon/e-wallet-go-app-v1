package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
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
	var req dto.TransferInquiryRequest
	if err := fctx.BodyParser(&req); err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	inquiryKey, err := t.transactionService.TransferInquiry(fctx.Context(), req)

	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return fctx.Status(200).JSON(dto.TransferInquiryResponse{
		Status:  true,
		Message: "success get inquiry key",
		Data:    inquiryKey,
	})
}

func (t *transferApi) transferExecute(fctx *fiber.Ctx) error {
	var req dto.TransferExecuteRequest
	if err := fctx.BodyParser(&req); err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err := t.transactionService.TranferExecute(fctx.Context(), req)
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return fctx.Status(200).JSON(dto.BasicResponse{
		Status:  true,
		Message: "success tranfer executed",
	})
}
