package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	repository "github.com/midoon/e-wallet-go-app-v1/Repository"
	service "github.com/midoon/e-wallet-go-app-v1/Service"
	"github.com/midoon/e-wallet-go-app-v1/api"
	"github.com/midoon/e-wallet-go-app-v1/internal/component"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/midoon/e-wallet-go-app-v1/middleware"
)

func main() {

	cnf := config.GetConfig()
	validator := validator.New()
	dbConnection := component.GetDBOpenConenction(cnf)

	userRepository := repository.NewUserRepository(dbConnection)
	tokenRepository := repository.NewTokenRepository(dbConnection)
	accountRepository := repository.NewAccountRepository(dbConnection)

	userService := service.NewUserService(userRepository, tokenRepository, accountRepository, validator, cnf)

	authMidd := middleware.AuthMiddleware(cnf)

	app := fiber.New()
	api.NewAuthApi(app, userService, authMidd)

	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		panic(err)
	}
}
