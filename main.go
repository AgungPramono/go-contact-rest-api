package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/controller"
	"go-contact-rest-api/helper"
	"go-contact-rest-api/repository"
	"go-contact-rest-api/service/impl"
)

func main() {

	validate := validator.New()

	app := fiber.New()
	db, _ := helper.ConnectDB()

	userRepository := repository.NewUserRepository()
	userService := impl.NewUserService(userRepository, db, validate)
	authService := impl.NewAuthService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	handler := helper.Handler{
		UserController: *userController,
		AuthController: *authController,
	}

	helper.SetupRouter(app, &handler, userService)
	app.Listen(":8080")
}
