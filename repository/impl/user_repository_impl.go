package impl

import (
	"go-contact-rest-api/model"
	"go-contact-rest-api/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(user *model.User, db *gorm.DB) error {
	return db.Create(user).Error
}

func (repository *UserRepositoryImpl) Update(user *model.User, db *gorm.DB) error {
	return db.Save(user).Error
}

func (repository *UserRepositoryImpl) Delete(id string, db *gorm.DB) error {
	return db.Delete(&model.User{}, "id = ?", id).Error
}

func (repository *UserRepositoryImpl) ExistById(userName string, db *gorm.DB) (*model.User, error) {
	var user model.User
	err := db.First(&user, "username = ?", userName).Error
	return &user, err
}

func (repository *UserRepositoryImpl) FindAll(db *gorm.DB) ([]*model.User, error) {
	var users []*model.User
	err := db.Find(&users).Error
	return users, err
}

func (repository *UserRepositoryImpl) FindById(id string, db *gorm.DB) (*model.User, error) {
	var user model.User
	err := db.First(&user, "id = ?", id).Error
	return &user, err
}

func (repository *UserRepositoryImpl) FindByToken(token string, db *gorm.DB) (*model.User, error) {
	var user model.User
	err := db.First(&user, "token = ?", token).Error
	return &user, err
}

func (repository *UserRepositoryImpl) FindByUsername(username string, db *gorm.DB) (*model.User, error) {
	var user model.User
	err := db.First(&user, "username = ?", username).Error
	return &user, err
}
