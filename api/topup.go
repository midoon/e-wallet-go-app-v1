package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
)

type topUpApi struct {
	topupService domain.TopupService
}

func NewTopUpApi(app *fiber.App, authMid fiber.Handler, topupService domain.TopupService) {
	t := topUpApi{
		topupService: topupService,
	}

	app.Post("/api/topup", authMid, t.InitializeTopUp)

}

func (t *topUpApi) InitializeTopUp(fctx *fiber.Ctx) error {
	var req dto.TopUpRequest
	if err := fctx.BodyParser(&req); err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	req.UserId = fctx.Locals("x-user-id").(string)

	res, err := t.topupService.InitializeTopUp(fctx.Context(), req)
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return fctx.Status(200).JSON(res)
}
