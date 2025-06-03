package impl

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/app"
	"go-contact-rest-api/model"
	"go-contact-rest-api/repository"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web/request"
	"go-contact-rest-api/web/response"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func (service UserServiceImpl) FindByToken(token string) (*model.User, error) {
	return service.repository.FindByToken(token, service.DB)
}

func NewUserService(repository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) service.UserService {
	return &UserServiceImpl{
		repository: repository,
		DB:         DB,
		Validate:   validate,
	}
}

func (service UserServiceImpl) Register(request *request.RegisterUserRequest) error {

	err2 := service.Validate.Struct(request)
	if err2 != nil {
		parseErr := app.ParseValidationErrors(err2)
		return fiber.NewError(fiber.StatusBadRequest, parseErr)
	}

	_, err := service.repository.ExistById(request.Username, service.DB)
	if err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Username already exist")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	user := model.User{
		Name:     request.Name,
		Username: request.Username,
		Password: string(password),
		Token:    nil,
	}
	return service.repository.Save(&user, service.DB)
}

func Get(user *model.User) response.UserResponse {
	return response.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}
}

func (service UserServiceImpl) Update(user *model.User, request *request.UserUpdateRequest) (response.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return response.UserResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		}
		user.Password = string(password)
	}

	err = service.repository.Update(user, service.DB)
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update user")
	}
	return Get(user), nil
}

func (service UserServiceImpl) Get(user *model.User) (response.UserResponse, error) {
	return Get(user), nil
}
