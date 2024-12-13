package service

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_errors.RestError {
	logger.Info(
		"Init updateUser model.",
		zap.String("journey", "updateUser"),
	)

	err := ud.userRepository.(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"),
		)
		return nil, err
	}

	logger.Info("updateUser model executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	return nil
}
