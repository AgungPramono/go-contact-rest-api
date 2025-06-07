package controller

import (
	"github.com/gofiber/fiber/v2"
	error2 "go-contact-rest-api/helper/error"
	"go-contact-rest-api/model"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
)

type AddressController struct {
	addressService service.AddressService
}

func NewAddressController(service service.AddressService) *AddressController {
	return &AddressController{
		addressService: service,
	}
}

func (controller *AddressController) Create(ctx *fiber.Ctx) error {
	addressRequest := request.CreateAddressRequest{}
	err := ctx.BodyParser(&addressRequest)
	if err != nil {
		return error2.ResponseError(ctx, err, "error parsing request")
	}

	addressRequest.ContactID = ctx.Params("idContact")

	user := ctx.Locals("user").(*model.User)
	address, err := controller.addressService.Create(user, addressRequest)
	if err != nil {
		statusResponse := web.StatusResponse{
			Errors: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}
	response := web.ApiResponse{
		Data: address,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *AddressController) Update(ctx *fiber.Ctx) error {
	addressRequest := request.UpdateAddressRequest{}
	err := ctx.BodyParser(&addressRequest)
	if err != nil {
		return error2.ResponseError(ctx, err, "error parsing request")
	}

	addressRequest.AddressID = ctx.Params("idAddress")
	addressRequest.ContactID = ctx.Params("idContact")

	address, err := controller.addressService.Update(ctx.Locals("user").(*model.User), addressRequest)
	if err != nil {
		statusResponse := web.StatusResponse{
			Errors: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}
	response := web.ApiResponse{
		Data: address,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *AddressController) Delete(ctx *fiber.Ctx) error {
	err := controller.addressService.Delete(ctx.Locals("user").(*model.User), ctx.Params("idContact"), ctx.Params("idAddress"))
	statusResponse := web.StatusResponse{}
	if err != nil {
		statusResponse.Errors = err.Error()
		return ctx.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}
	statusResponse.Data = "OK"
	return ctx.Status(fiber.StatusOK).JSON(statusResponse)
}

func (controller *AddressController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	addressResponse, err := controller.addressService.Get(user, ctx.Params("idContact"), ctx.Params("idAddress"))
	if err != nil {
		statusResponse := web.StatusResponse{
			Errors: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}
	response := web.ApiResponse{
		Data: addressResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *AddressController) ListAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)

	addresses, err := controller.addressService.FindAll(user, ctx.Params("idContact"))
	if err != nil {
		statusResponse := web.StatusResponse{
			Errors: err.Error(),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(statusResponse)
	}
	response := web.ApiResponse{
		Data: addresses,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
