package service

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_errors.RestError) {

	logger.Info("Init loginUser model", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, err
	}

	logger.Info("loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
	)
	return user, nil
}
