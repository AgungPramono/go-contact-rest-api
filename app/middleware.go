package app

import (
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/service"
	"time"
)

func UserResolver(userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("X-API-TOKEN")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: missing token",
			})
		}

		user, err := userService.FindByToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: invalid token",
			})
		}

		if user.TokenExpiredAt < time.Now().Unix() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token expired",
			})
		}

		// Set user to context
		c.Locals("user", user)
		return c.Next()
	}
}
