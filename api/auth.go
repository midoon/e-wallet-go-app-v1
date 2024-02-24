package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
)

type authApi struct {
	userService domain.UserService
}

func NewAuthApi(app *fiber.App, userService domain.UserService) {
	handler := authApi{
		userService: userService,
	}
	app.Post("/api/register", handler.register)
}

func (auth *authApi) register(ctx *fiber.Ctx) error {
	var req dto.UserRegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	res, err := auth.userService.Register(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(helper.HttpStatusErr(err))
	}

	return ctx.Status(200).JSON(res)
}
