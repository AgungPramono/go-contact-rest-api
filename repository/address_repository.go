package repository

import (
	"go-contact-rest-api/model"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Save(address *model.Address, db *gorm.DB) error
	Update(address *model.Address, db *gorm.DB) error
	Delete(id string, db *gorm.DB) error
	FindFirstByContactAndId(contactId string, addressId string, db *gorm.DB) (*model.Address, error)
	FindAllByContactId(contactId string, db *gorm.DB) ([]*model.Address, error)
}
