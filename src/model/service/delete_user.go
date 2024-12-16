package service

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string) *rest_errors.RestError {
	logger.Info(
		"Init deleteUser model.",
		zap.String("journey", "deleteUser"),
	)

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"),
		)
		return err
	}

	logger.Info("deleteUser model executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	return nil
}
