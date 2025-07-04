package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	app2 "go-contact-rest-api/app"
	"go-contact-rest-api/controller"
	"go-contact-rest-api/service"
)

type Handler struct {
	UserController    controller.UserController
	AuthController    controller.AuthController
	ContactController controller.ContactController
	AddressController controller.AddressController
}

func SetupRouter(app *fiber.App, h *Handler, service service.UserService) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081",
		AllowHeaders: "X-Api-Token,Origin, Content-Type, Accept,Authorization,Access-Control-Allow-Header,Access-Control-Allow-Origin",
		AllowMethods: "GET, POST, PUT, DELETE,PATCH",
	}))

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
	api.Get("/contacts", h.ContactController.SearchContact)

	api.Get("/contacts/:idContact/addresses", h.AddressController.ListAll)
	api.Post("/contacts/:idContact/addresses", h.AddressController.Create)
	api.Put("/contacts/:idContact/addresses/:idAddress", h.AddressController.Update)
	api.Delete("/contacts/:idContact/addresses/:idAddress", h.AddressController.Delete)
	api.Get("/contacts/:idContact/addresses/:idAddress", h.AddressController.Get)
}
