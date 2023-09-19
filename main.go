package main

import (
	"errors"
	"fmt"
	"log"
	_ "nautilus/config"
	"nautilus/database"
	"nautilus/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			fmt.Println(err.Error())

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}

			// Return from handler
			return nil
		},
	})
	app.Use(cors.New())
	app.Use(recover.New())

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
