package controller

import (
	"net/http"
	"strings"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/validation"
	"github.com/Michael-Andryeer/crud-go/src/controller/model/request"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser Controller",
		zap.String("journey", "UpdateUser"),
		zap.String("operation", "init"),
	)

	var userRequest request.UserUpdateRequest
	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" { // Se mandar um Id vazio da erro
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error(
			"Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("updateUser controller executed successfully",
		zap.String("journey", userId),
		zap.String("journey", "updateUser"),
	)
	c.Status(http.StatusOK)
}
