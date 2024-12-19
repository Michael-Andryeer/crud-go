package service

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_errors.RestError) {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		return nil, rest_errors.NewBadRequestError("Email already exists")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call CreateUser repository",
			err,
			zap.String("journey", "createUser"),
		)
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("journey", "createUser"),
	)
	return userDomainRepository, nil
}
