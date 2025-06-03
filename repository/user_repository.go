package repository

import (
	"go-contact-rest-api/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *model.User, db *gorm.DB) error
	Update(user *model.User, db *gorm.DB) error
	Delete(id string, db *gorm.DB) error
	ExistById(userName string, db *gorm.DB) (*model.User, error)
	FindAll(db *gorm.DB) ([]*model.User, error)
	FindById(id string, db *gorm.DB) (*model.User, error)
	FindByToken(token string, db *gorm.DB) (*model.User, error)
	FindByUsername(username string, db *gorm.DB) (*model.User, error)
}
