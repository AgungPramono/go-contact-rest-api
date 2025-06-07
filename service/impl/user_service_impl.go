package impl

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-contact-rest-api/app"
	"go-contact-rest-api/model"
	"go-contact-rest-api/repository"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
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

func (service UserServiceImpl) Register(request *request.RegisterUserRequest) web.StatusResponse {

	err2 := service.Validate.Struct(request)
	if err2 != nil {
		parseErr := app.ParseValidationErrors(err2)
		return web.CreateResponse(fiber.StatusBadRequest, parseErr, nil, false)
	}

	_, err := service.repository.ExistById(request.Username, service.DB)
	if err == nil {
		return web.CreateResponse(fiber.StatusBadRequest, "Username already exist", nil, false)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CreateResponse(fiber.StatusInternalServerError, "Internal server error", nil, false)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.CreateResponse(fiber.StatusInternalServerError, "Internal server error", nil, false)
	}

	user := model.User{
		Name:     request.Name,
		Username: request.Username,
		Password: string(password),
		Token:    nil,
	}
	err = service.repository.Save(&user, service.DB)
	if err != nil {
		return web.CreateResponse(fiber.StatusInternalServerError, "register user failed", nil, false)
	}
	return web.CreateResponse(fiber.StatusOK, nil, "OK", true)
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
