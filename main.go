package main

import (
	"log"
	_ "nautilus/config"
	"nautilus/database"
	"nautilus/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	database.ConnectDB()
}

func main() {
	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "Nautilus",
		AppName:      "Nautilus App v1.0.1",
	})
	app.Use(cors.New())

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
