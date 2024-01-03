package routes

import "github.com/gofiber/fiber/v2"

func RegisterRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 NotFound")
	})
}
