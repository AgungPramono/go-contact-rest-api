package service

import (
	"go-contact-rest-api/model"
	"go-contact-rest-api/web"
	"go-contact-rest-api/web/request"
	"go-contact-rest-api/web/response"
)

type UserService interface {
	Register(request *request.RegisterUserRequest) web.StatusResponse
	Update(user *model.User, request *request.UserUpdateRequest) (response.UserResponse, error)
	FindByToken(token string) (*model.User, error)
	Get(user *model.User) (response.UserResponse, error)
}
