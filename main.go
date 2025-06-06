package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/config"
	"go-contact-rest-api/controller"
	"go-contact-rest-api/helper"
	impl2 "go-contact-rest-api/repository/impl"
	"go-contact-rest-api/service/impl"
)

func main() {

	validate := validator.New()

	app := fiber.New()
	db, _ := config.ConnectDB()

	userRepository := impl2.NewUserRepository()
	contactRepository := impl2.NewContactRepositoryImpl()
	addressRepository := impl2.NewAddressRepositoryImpl()

	userService := impl.NewUserService(userRepository, db, validate)
	authService := impl.NewAuthService(userRepository, db, validate)
	addressService := impl.NewAddressService(addressRepository, contactRepository, db, validate)
	contactService := impl.NewContactService(contactRepository, db, validate)

	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)
	contactController := controller.NewContactController(contactService)
	addressController := controller.NewAddressController(addressService)

	handler := helper.Handler{
		UserController:    *userController,
		AuthController:    *authController,
		ContactController: *contactController,
		AddressController: *addressController,
	}

	helper.SetupRouter(app, &handler, userService)

	app.Listen(":8080")
}
