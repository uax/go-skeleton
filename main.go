package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uax/go-skeleton/config"
	"github.com/uax/go-skeleton/routes"
)

func main() {
	config.InitDatabase()
	app := fiber.New()
	routes.RegisterRouter(app)
	app.Listen(config.Config("APP_LISTEN"))
}
