package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/api"
	"github.com/midoon/e-wallet-go-app-v1/internal/component"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/midoon/e-wallet-go-app-v1/middleware"
	"github.com/midoon/e-wallet-go-app-v1/repository"
	"github.com/midoon/e-wallet-go-app-v1/service"
)

func main() {

	cnf := config.GetConfig()
	validator := validator.New()
	dbConnection := component.GetDBOpenConenction(cnf)
	rdbConnection := component.GetRedisConnection(cnf)
	rmqConnection := component.GetRabbitMQConnection(cnf)

	userRepository := repository.NewUserRepository(dbConnection)
	tokenRepository := repository.NewTokenRepository(dbConnection)
	accountRepository := repository.NewAccountRepository(dbConnection)
	transactionRepository := repository.NewTransactionRepository(dbConnection)
	notificationRepository := repository.NewNotificationRepository(dbConnection)

	userService := service.NewUserService(userRepository, tokenRepository, accountRepository, validator, cnf)
	transactionService := service.NewTransactionService(transactionRepository, accountRepository, notificationRepository, rdbConnection, validator, rmqConnection, cnf)
	notificationService := service.NewNotificationService(notificationRepository, accountRepository)

	authMidd := middleware.AuthMiddleware(cnf)

	app := fiber.New()
	api.NewAuthApi(app, userService, authMidd)
	api.NewTranferApi(app, transactionService, authMidd)
	api.NewNotificationApi(notificationService, app, authMidd)

	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		panic(err)
	}
}
