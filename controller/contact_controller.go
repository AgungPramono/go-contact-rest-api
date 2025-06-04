package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/model"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
)

type ContactController struct {
	ContactService service.ContactService
}

func NewContactController(contactService service.ContactService) *ContactController {
	return &ContactController{
		ContactService: contactService,
	}
}

func (controller *ContactController) Create(c *fiber.Ctx) error {
	create := &request.CreateContactRequest{}
	if err := c.BodyParser(&create); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	user := c.Locals("user").(*model.User)
	contactResponse, err := controller.ContactService.Create(user, create)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	response := web.ApiResponse{
		Data: contactResponse,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *ContactController) GetContact(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)
	contactResponse, err := controller.ContactService.Get(user, c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	response := web.ApiResponse{
		Data: contactResponse,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
