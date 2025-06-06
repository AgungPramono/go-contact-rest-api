package repository

import (
	"go-contact-rest-api/model"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Save(address *model.Address, db *gorm.DB) error
	Update(address *model.Address, db *gorm.DB) error
	Delete(id string, db *gorm.DB) error
	FindAll(user *model.User, db *gorm.DB) ([]*model.Address, error)
	FindById(id string, db *gorm.DB) (*model.Address, error)
	FindByUserId(user *model.User, id string, db *gorm.DB) (*model.Address, error)
	FindFirstByContactAndId(contactId string, addressId string, db *gorm.DB) (*model.Address, error)
	FindAllByContactId(contactId string, db *gorm.DB) ([]*model.Address, error)
}
