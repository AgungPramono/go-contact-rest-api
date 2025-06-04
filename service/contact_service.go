package service

import (
	"go-contact-rest-api/model"
	"go-contact-rest-api/web/request"
	"go-contact-rest-api/web/response"
)

type ContactService interface {
	Create(user *model.User, request *request.CreateContactRequest) (response.ContactResponse, error)
	Get(user *model.User, id string) (response.ContactResponse, error)
	Update(user *model.User, request *request.UpdateContactRequest) (response.ContactResponse, error)
	Delete(user *model.User, id string) error
}
