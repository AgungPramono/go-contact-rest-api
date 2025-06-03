package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/model"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (usrController *UserController) Register(ctx *fiber.Ctx) error {

	registerUserRequest := &request.RegisterUserRequest{}
	err := ctx.BodyParser(registerUserRequest)
	if err != nil {
		response := web.ApiResponse{
			Status:  false,
			Message: "invalid request",
			Errors:  err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = usrController.UserService.Register(registerUserRequest)
	if err != nil {
		response := web.ApiResponse{
			Status:  false,
			Message: "Register failed",
			Errors:  err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := web.ApiResponse{
		Data: "OK",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (usrController *UserController) Update(ctx *fiber.Ctx) error {
	userUpdateRequest := &request.UserUpdateRequest{}
	err := ctx.BodyParser(userUpdateRequest)
	if err != nil {
		response := web.ApiResponse{
			Status:  false,
			Message: "invalid request",
			Errors:  err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	userVal := ctx.Locals("user").(*model.User)
	data, err := usrController.UserService.Update(userVal, userUpdateRequest)
	if err != nil {
		response := web.ApiResponse{
			Status:  false,
			Message: "Register failed",
			Errors:  err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := web.ApiResponse{
		Data: data,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (usrController *UserController) Get(ctx *fiber.Ctx) error {
	userResponse, _ := usrController.UserService.Get(ctx.Locals("user").(*model.User))
	response := web.ApiResponse{
		Data:   userResponse,
		Status: true,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
