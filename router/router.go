package router

import (
	"nautilus/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", func(c *fiber.Ctx) error {
		panic("This panic is caught by fiber")
	})

	api.Get("/login", handler.Login)
}
