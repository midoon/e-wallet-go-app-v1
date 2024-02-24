package main

import (
	"github.com/gofiber/fiber/v2"
	repository "github.com/midoon/e-wallet-go-app-v1/Repository"
	service "github.com/midoon/e-wallet-go-app-v1/Service"
	"github.com/midoon/e-wallet-go-app-v1/internal/api"
	"github.com/midoon/e-wallet-go-app-v1/internal/component"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
)

func main() {

	cnf := config.GetConfig()
	dbConnection := component.GetDBOpenConenction(cnf)
	userRepository := repository.NewUserRepository(dbConnection)
	userService := service.NewUserService(userRepository)

	app := fiber.New()
	api.NewAuthApi(app, userService)
	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		panic(err)
	}
}
