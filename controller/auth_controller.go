package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/model"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ac *AuthController) Login(ctx *fiber.Ctx) error {
	loginRequest := request.LoginRequest{}
	err := ctx.BodyParser(&loginRequest)
	if err != nil {
		response := web.ApiResponse{
			Status:  false,
			Message: "invalid request",
			Errors:  err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	tokenResponse, err := ac.AuthService.Login(&loginRequest)
	if err != nil {
		apiResponse := web.ApiResponse{
			Errors: err.Error(),
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(apiResponse)
	}
	response := web.ApiResponse{
		Data: tokenResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (ac *AuthController) Logout(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	err := ac.AuthService.Logout(user)
	if err != nil {
		apiResponse := web.ApiResponse{
			Status: false,
			Errors: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}
	response := web.ApiResponse{
		Data: "OK",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
