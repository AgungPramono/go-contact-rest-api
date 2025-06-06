package impl

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-contact-rest-api/app"
	"go-contact-rest-api/model"
	"go-contact-rest-api/repository"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
	"go-contact-rest-api/web/response"
	"gorm.io/gorm"
	"math"
)

type ContactServiceImpl struct {
	repository repository.ContactRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewContactService(contactRepository repository.ContactRepository, db *gorm.DB, validate *validator.Validate) service.ContactService {
	return &ContactServiceImpl{
		repository: contactRepository,
		DB:         db,
		Validate:   validate,
	}
}

func (contactService ContactServiceImpl) Create(user *model.User, request *request.CreateContactRequest) (response.ContactResponse, error) {
	err := contactService.Validate.Struct(request)
	if err != nil {
		ParseErr := app.ParseValidationErrors(err)
		return response.ContactResponse{}, fiber.NewError(fiber.StatusBadRequest, ParseErr)
	}

	contact := model.Contact{
		ID:        uuid.New().String(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Username:  user.Username,
	}

	err = contactService.repository.Save(&contact, contactService.DB)
	if err != nil {
		return response.ContactResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to create contact")
	}

	return toContactResponse(contact)
}

func toContactResponse(contact model.Contact) (response.ContactResponse, error) {
	return response.ContactResponse{
		Id:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}, nil
}

func (contactService ContactServiceImpl) Get(user *model.User, id string) (response.ContactResponse, error) {

	contact, err := contactService.repository.FindFirstByUserId(*user, id, contactService.DB)
	if err != nil {
		return response.ContactResponse{}, fiber.NewError(fiber.StatusNotFound, "Contact not found")
	}
	return toContactResponse(*contact)
}

func (contactService ContactServiceImpl) Update(user *model.User, request *request.UpdateContactRequest) (response.ContactResponse, error) {
	err := contactService.Validate.Struct(request)
	if err != nil {
		ParseErr := app.ParseValidationErrors(err)
		return response.ContactResponse{}, fiber.NewError(fiber.StatusBadRequest, ParseErr)
	}

	contact, err := contactService.repository.FindFirstByUserId(*user, request.Id, contactService.DB)
	if err != nil {
		return response.ContactResponse{}, fiber.NewError(fiber.StatusNotFound, "Contact not found")
	}

	contact.FirstName = request.FirstName
	contact.LastName = request.LastName
	contact.Phone = request.Phone
	contact.Email = request.Email
	err = contactService.repository.Update(contact, contactService.DB)
	if err != nil {
		return response.ContactResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update contact")
	}

	return toContactResponse(*contact)
}

func (contactService ContactServiceImpl) Delete(user *model.User, id string) error {
	_, err := contactService.repository.FindFirstByUserId(*user, id, contactService.DB)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Contact not found")
	}
	return contactService.repository.Delete(id, contactService.DB)
}

func (contactService ContactServiceImpl) SearchContacts(user *model.User, search request.SearchContactRequest) ([]response.ContactResponse, web.PagingResponse, error) {
	contactResult, total, err := contactService.repository.SearchContacts(
		user,
		search,
		contactService.DB)

	if err != nil {
		return nil, web.PagingResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	var data []response.ContactResponse
	for _, contact := range contactResult {
		contactData := response.ContactResponse{
			Id:        contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Phone:     contact.Phone,
			Email:     contact.Email,
		}
		data = append(data, contactData)
	}

	totalPage := int(math.Ceil(float64(total) / float64(search.Size)))

	Paging := web.PagingResponse{
		CurrentPage: search.Page,
		TotalPage:   totalPage,
		Size:        search.Size,
	}

	return data, Paging, nil
}
