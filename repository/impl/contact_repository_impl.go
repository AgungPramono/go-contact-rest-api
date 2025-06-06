package impl

import (
	"go-contact-rest-api/model"
	"go-contact-rest-api/web/request"
	"gorm.io/gorm"
)

type ContactRepositoryImpl struct {
}

const DefaultPageSize = 10

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

func (r *ContactRepositoryImpl) SearchContacts(user *model.User, search request.SearchContactRequest, db *gorm.DB) ([]model.Contact, int64, error) {
	var contacts []model.Contact
	var total int64

	query := db.Model(&model.Contact{})

	query = query.Where("username = ?", user.Username)

	if search.Name != "" {
		query = query.Where("first_name LIKE ? OR last_name LIKE ?", "%"+search.Name+"%", "%"+search.Name+"%")
	}

	if search.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+search.Phone+"%")
	}

	if search.Email != "" {
		query = query.Where("email LIKE ?", "%"+search.Email+"%")
	}

	// Count the total number of results before applying pagination
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (search.Page - 1) * search.Size

	err = query.Offset(offset).Limit(search.Size).Find(&contacts).Error
	if err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}
