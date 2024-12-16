package controller

import (
	"net/http"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/configuration/validation"
	"github.com/Michael-Andryeer/crud-go/src/controller/model/request"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser Controller",
		zap.String("journey", "UpdateUser"),
		zap.String("operation", "init"),
	)

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil { // Se mandar um Id vazio da erro
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_errors.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
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
