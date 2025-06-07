package impl

import (
	"go-contact-rest-api/model"
	"gorm.io/gorm"
)

type AddressRepositoryImpl struct{}

func NewAddressRepositoryImpl() *AddressRepositoryImpl {
	return &AddressRepositoryImpl{}
}

func (r *AddressRepositoryImpl) Save(address *model.Address, db *gorm.DB) error {
	return db.Create(address).Error
}

func (r *AddressRepositoryImpl) Update(address *model.Address, db *gorm.DB) error {
	return db.Save(address).Error
}

func (r *AddressRepositoryImpl) Delete(id string, db *gorm.DB) error {
	return db.Delete(&model.Address{}, "id = ?", id).Error
}

func (r *AddressRepositoryImpl) FindAllByContactId(contactId string, db *gorm.DB) ([]*model.Address, error) {
	var addresses []*model.Address
	err := db.Where("contact_id = ?", contactId).Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (r *AddressRepositoryImpl) FindFirstByContactAndId(contactId string, addressId string, db *gorm.DB) (*model.Address, error) {
	var address model.Address
	err := db.Where("id = ? and contact_id = ?", addressId, contactId).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}
