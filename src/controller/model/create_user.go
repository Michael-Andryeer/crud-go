package model

import (
	"fmt"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_errors.RestError {
	logger.Info("Init CreateUser method", zap.String("journey", "createUser"))

	ud.EncryptPassowrd()
	fmt.Println(ud)
	return nil
}
