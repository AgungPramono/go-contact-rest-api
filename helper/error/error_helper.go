package error

import (
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/web"
)

func ResponseError(ctx *fiber.Ctx, err error, message string) error {
	if err != nil {
		response := web.ApiResponse{
			Status:  false,
			Message: message,
			Errors:  err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)

	}
	return nil
}
