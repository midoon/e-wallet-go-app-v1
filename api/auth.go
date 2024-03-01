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

func NewAuthApi(app *fiber.App, userService domain.UserService, authMidd fiber.Handler) {
	handler := authApi{
		userService: userService,
	}
	app.Post("/api/auth/register", handler.register)
	app.Post("/api/auth/login", handler.login)
	app.Delete("/api/auth/logout", authMidd, handler.logout)
}

func (auth *authApi) register(fctx *fiber.Ctx) error {
	var req dto.UserRegisterRequest
	if err := fctx.BodyParser(&req); err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err := auth.userService.Register(fctx.Context(), req)
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	res := dto.UserRegisterResponse{
		Status:  true,
		Message: "success register",
	}

	return fctx.Status(200).JSON(res)
}

func (auth *authApi) login(fctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := fctx.BodyParser(&req); err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	tokenData, err := auth.userService.Login(fctx.Context(), req)
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	res := dto.LoginResponse{
		Status:  true,
		Message: "success login",
		Data:    tokenData,
	}
	return fctx.Status(200).JSON(res)
}

func (auth *authApi) logout(fctx *fiber.Ctx) error {
	userId := fctx.Locals("x-user-id").(string)
	err := auth.userService.Logout(fctx.Context(), string(userId))
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	res := dto.LogoutResponse{
		Status:  true,
		Message: "success logout",
	}
	return fctx.Status(200).JSON(res)
}
