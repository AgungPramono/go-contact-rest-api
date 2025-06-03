package service

import (
	"go-contact-rest-api/model"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
)

type AuthService interface {
	Login(loginUser *request.LoginRequest) (web.TokenResponse, error)
	Logout(user *model.User) error
}
