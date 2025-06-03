package impl

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-contact-rest-api/app"
	"go-contact-rest-api/model"
	"go-contact-rest-api/repository"
	"go-contact-rest-api/service"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthServiceImpl struct {
	repository repository.UserRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewAuthService(repository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) service.AuthService {
	return &AuthServiceImpl{
		repository: repository,
		DB:         DB,
		Validate:   validate,
	}
}

func (service AuthServiceImpl) Login(loginUser *request.LoginRequest) (web.TokenResponse, error) {
	err := service.Validate.Struct(loginUser)
	if err != nil {
		parseErr := app.ParseValidationErrors(err)
		return web.TokenResponse{}, fiber.NewError(fiber.StatusBadRequest, parseErr)
	}

	user, err := service.repository.FindByUsername(loginUser.Username, service.DB)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.TokenResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Username or password is wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		return web.TokenResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Username or password is wrong")
	}
	tokenUUID := uuid.New().String()
	user.Token = &tokenUUID
	next30day := time.Now().Add(30 * 24 * time.Hour).Unix()
	user.TokenExpiredAt = next30day

	err = service.repository.Update(user, service.DB)
	if err != nil {
		return web.TokenResponse{}, err
	}

	return web.TokenResponse{
		Token:     tokenUUID,
		ExpiredAt: &next30day,
	}, nil
}

func (service AuthServiceImpl) Logout(user *model.User) error {
	user.Token = nil
	user.TokenExpiredAt = 0
	return service.repository.Update(user, service.DB)
}
