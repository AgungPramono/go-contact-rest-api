package repository

import (
	"go-contact-rest-api/model"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Save(contact *model.Contact, db *gorm.DB) error
	FindFirstByUserId(user model.User, id string, db *gorm.DB) (*model.Contact, error)
	Update(contact *model.Contact, db *gorm.DB) error
	Delete(id string, db *gorm.DB) error
}
