package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/model"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
	"strconv"
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
		return GetResponse(c, err)
	}

	user := c.Locals("user").(*model.User)
	contactResponse, err := controller.ContactService.Create(user, create)
	if err != nil {
		apiResponse := web.ApiResponse{
			Errors: err.Error(),
			Status: false,
		}
		return c.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	response := web.ApiResponse{
		Data: contactResponse,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetResponse(c *fiber.Ctx, err error) error {
	apiResponse := web.ApiResponse{
		Status: false,
		Errors: err.Error(),
	}
	return c.Status(fiber.StatusBadRequest).JSON(apiResponse)
}

func (controller *ContactController) GetContact(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)
	contactResponse, err := controller.ContactService.Get(user, c.Params("id"))
	if err != nil {
		statusResponse := web.StatusResponse{
			Errors: err.Error(),
		}
		return c.Status(fiber.StatusNotFound).JSON(statusResponse)
	}

	response := web.ApiResponse{
		Data: contactResponse,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *ContactController) Update(c *fiber.Ctx) error {
	updateRequest := &request.UpdateContactRequest{}
	if err := c.BodyParser(&updateRequest); err != nil {
		statusResponse := web.StatusResponse{
			Errors: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}

	Id := c.Params("id")
	updateRequest.Id = Id
	user := c.Locals("user").(*model.User)
	contactResponse, err := controller.ContactService.Update(user, updateRequest)
	if err != nil {
		statusResponse := web.StatusResponse{
			Errors: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}

	response := web.ApiResponse{
		Data: contactResponse,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *ContactController) Delete(c *fiber.Ctx) error {
	Id := c.Params("id")
	user := c.Locals("user").(*model.User)

	statusResponse := web.StatusResponse{}
	
	err := controller.ContactService.Delete(user, Id)
	if err != nil {
		statusResponse.Errors = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}

	statusResponse.Data = "OK"
	return c.Status(fiber.StatusOK).JSON(statusResponse)
}

func (controller *ContactController) SearchContact(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 5
	}
	searchRequest := request.SearchContactRequest{
		Name:  c.Params("name"),
		Email: c.Params("email"),
		Phone: c.Params("phone"),
		Page:  page,
		Size:  size,
	}
	statusResponse := web.StatusResponse{}
	if err := c.QueryParser(&searchRequest); err != nil {
		statusResponse.Errors = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}

	contacts, paging, err := controller.ContactService.SearchContacts(user, searchRequest)
	if err != nil {
		statusResponse.Errors = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}
	response := web.ApiResponse{
		Data:   contacts,
		Paging: &paging,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
