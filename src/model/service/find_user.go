package service

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_errors.RestError) {

	logger.Info("Init FindUserByIDServices service",
		zap.String("journey", "FindUserByIDServices"),
	)
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_errors.RestError) {
	logger.Info("Init FindUserByEmailServices service",
		zap.String("journey", "FindUserByEmailServices"),
	)
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email, password string) (model.UserDomainInterface, *rest_errors.RestError) {
	logger.Info("Init FindUserByEmailAndPasswordServices service",
		zap.String("journey", "FindUserByEmailAndPasswordServices"),
	)
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
