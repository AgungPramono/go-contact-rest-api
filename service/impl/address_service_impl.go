package impl

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-contact-rest-api/app"
	"go-contact-rest-api/model"
	"go-contact-rest-api/repository"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web/request"
	"go-contact-rest-api/web/response"
	"gorm.io/gorm"
)

type AddressService struct {
	addressRepository repository.AddressRepository
	contactRepository repository.ContactRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewAddressService(addressRepository repository.AddressRepository, contactRepository repository.ContactRepository, DB *gorm.DB, validate *validator.Validate) service.AddressService {
	return &AddressService{
		addressRepository: addressRepository,
		contactRepository: contactRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service AddressService) Create(user *model.User, request request.CreateAddressRequest) (*response.AddressResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		parseErr := app.ParseValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, parseErr)
	}

	var addressResponse *response.AddressResponse

	// Wrap code with GORM transaction
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		// Find the contact
		contact, err := service.contactRepository.FindFirstByUserId(*user, request.ContactID, tx)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fiber.NewError(fiber.StatusNotFound, "Contact not found")
			}
			return err
		}

		// Create the address
		address := model.Address{
			ID:         uuid.New().String(),
			ContactID:  contact.ID,
			Street:     request.Street,
			City:       request.City,
			Province:   request.Province,
			Country:    request.Country,
			PostalCode: request.PostalCode,
		}

		// Save the address
		err = service.addressRepository.Save(&address, tx)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create address")
		}

		// Convert saved address to response
		addressResponse = toAddressResponse(&address)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return addressResponse, nil
}

func toAddressResponse(address *model.Address) *response.AddressResponse {
	return &response.AddressResponse{
		Id:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		Country:    address.Country,
		PostalCode: address.PostalCode,
	}
}

func (service AddressService) Update(user *model.User, request request.UpdateAddressRequest) (*response.AddressResponse, error) {
	ValidateErr := service.Validate.Struct(request)
	if ValidateErr != nil {
		parseErr := app.ParseValidationErrors(ValidateErr)
		return nil, fiber.NewError(fiber.StatusBadRequest, parseErr)
	}

	contact, err := service.contactRepository.FindFirstByUserId(*user, request.ContactID, service.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Contact not found")
		}
	}

	address, err := service.addressRepository.FindFirstByContactAndId(contact.ID, request.AddressID, service.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Address not found")
		}
	}

	address.Street = request.Street
	address.City = request.City
	address.Province = request.Province
	address.Country = request.Country
	address.PostalCode = request.PostalCode

	err = service.addressRepository.Update(address, service.DB)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update address")
	}

	return toAddressResponse(address), nil
}

func (service AddressService) Delete(user *model.User, contactId string, addressId string) error {

	contact, err := service.contactRepository.FindFirstByUserId(*user, contactId, service.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Contact not found")
		}
	}

	address, err := service.addressRepository.FindFirstByContactAndId(contact.ID, addressId, service.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Address not found")
		}
	}

	err = service.addressRepository.Delete(address.ID, service.DB)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete address")
	}
	return nil
}

func (service AddressService) Get(user *model.User, contactId string, addressId string) (*response.AddressResponse, error) {
	contact, err := service.contactRepository.FindFirstByUserId(*user, contactId, service.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Contact not found")
		}
	}

	address, err := service.addressRepository.FindFirstByContactAndId(contact.ID, addressId, service.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Address not found")
		}
	}

	return toAddressResponse(address), nil
}

func (service AddressService) FindAll(user *model.User, contactId string) (*[]response.AddressResponse, error) {
	contact, err := service.contactRepository.FindFirstByUserId(*user, contactId, service.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Contact not found")
		}
	}

	addresses, err := service.addressRepository.FindAllByContactId(contact.ID, service.DB)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get address")
	}

	var addressResponses []response.AddressResponse
	for _, address := range addresses {
		data := toAddressResponse(address)
		addressResponses = append(addressResponses, *data)
	}
	return &addressResponses, nil

}
