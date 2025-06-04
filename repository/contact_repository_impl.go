package repository

import (
	"go-contact-rest-api/model"
	"gorm.io/gorm"
)

type ContactRepositoryImpl struct {
}

func NewContactRepositoryImpl() *ContactRepositoryImpl {
	return &ContactRepositoryImpl{}
}

func (r *ContactRepositoryImpl) Save(contact *model.Contact, db *gorm.DB) error {
	return db.Create(contact).Error
}

func (r *ContactRepositoryImpl) FindFirstByUserId(user model.User, id string, db *gorm.DB) (*model.Contact, error) {
	var contact model.Contact
	err := db.Where("id = ? and username = ?", id, user.Username).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *ContactRepositoryImpl) Update(contact *model.Contact, db *gorm.DB) error {
	return db.Save(contact).Error
}

func (r *ContactRepositoryImpl) Delete(id string, db *gorm.DB) error {
	return db.Delete(&model.Contact{}, "id = ?", id).Error
}
