package helper

import (
	"github.com/gofiber/fiber/v2"
	app2 "go-contact-rest-api/app"
	"go-contact-rest-api/controller"
	"go-contact-rest-api/service"
)

type Handler struct {
	UserController    controller.UserController
	AuthController    controller.AuthController
	ContactController controller.ContactController
}

func SetupRouter(app *fiber.App, h *Handler, service service.UserService) {
	app.Post("/api/auth/login", h.AuthController.Login)
	app.Post("/api/users", h.UserController.Register)

	api := app.Group("/api", app2.UserResolver(service))
	api.Patch("/users/current", h.UserController.Update)
	api.Get("/users/current", h.UserController.Get)
	api.Delete("/auth/logout", h.AuthController.Logout)
	api.Post("/contacts", h.ContactController.Create)
	api.Get("/contacts/:id", h.ContactController.GetContact)
	api.Put("/contacts/:id", h.ContactController.Update)
	api.Delete("/contacts/:id", h.ContactController.Delete)
}
