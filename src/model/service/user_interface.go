package service

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_errors.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestError
	DeleteUser(string) *rest_errors.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_errors.RestError)
}
