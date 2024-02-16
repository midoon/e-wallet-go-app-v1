package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	err := app.Listen("localhost:5000")
	if err != nil {
		panic(err)
	}
}
