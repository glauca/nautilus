package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// Handler
func Index(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Login(c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "nautilus",
			"exp": time.Now().AddDate(0, 1, 0).Unix(),
			"nbf": time.Now().Unix(),
			"iat": time.Now().Unix(),
		})

	key := viper.GetString("jwt.secret")
	t, err := token.SignedString([]byte(key))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
