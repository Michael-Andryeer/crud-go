package service

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"github.com/Michael-Andryeer/crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestError)
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestError
	DeleteUser(string) *rest_errors.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_errors.RestError)
}
