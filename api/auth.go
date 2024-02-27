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
	app.Post("/api/login", handler.login)
}

func (auth *authApi) register(ctx *fiber.Ctx) error {
	var req dto.UserRegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err := auth.userService.Register(ctx.Context(), req)
	if err != nil {
		return ctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	res := dto.UserRegisterResponse{
		Status:  true,
		Message: "success register",
	}

	return ctx.Status(200).JSON(res)
}

func (auth *authApi) login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	err := auth.userService.Login(ctx.Context(), req)
	if err != nil {
		return ctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	return ctx.SendStatus(200)
}
