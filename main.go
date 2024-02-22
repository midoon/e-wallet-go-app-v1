package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/midoon/e-wallet-go-app-v1/internal/api"
	"github.com/midoon/e-wallet-go-app-v1/internal/component"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
)

func main() {

	cnf := config.GetConfig()
	dbConnection := component.GetDBOpenConenction(cnf)

	fmt.Println(dbConnection)

	app := fiber.New()

	api.NewResgisterUserApi(app)

	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		panic(err)
	}
}
