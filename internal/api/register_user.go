package api

import "github.com/gofiber/fiber/v2"

type registerUserApi struct {
}

func NewResgisterUserApi(app *fiber.App) {
	handler := registerUserApi{}
	app.Post("/api/register", handler.register)
}

func (r *registerUserApi) register(ctx *fiber.Ctx) error {
	return ctx.SendString("ss")
}
