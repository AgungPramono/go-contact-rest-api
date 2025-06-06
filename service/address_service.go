package service

import (
	"go-contact-rest-api/model"
	"go-contact-rest-api/web/request"
	"go-contact-rest-api/web/response"
)

type AddressService interface {
	Create(user *model.User, request request.CreateAddressRequest) (*response.AddressResponse, error)
	Update(user *model.User, request request.UpdateAddressRequest) (*response.AddressResponse, error)
	Delete(user *model.User, contactId string, addressId string) error
	Get(user *model.User, contactId string, addressId string) (*response.AddressResponse, error)
	FindAll(user *model.User, contactId string) (*[]response.AddressResponse, error)
}
