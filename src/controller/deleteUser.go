package controller

import (
	"net/http"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser Controller",
		zap.String("journey", "DeleteUser"),
		zap.String("operation", "init"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_errors.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error(
			"Error trying to call deleteUser service",
			err,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteUser controller executed successfully",
		zap.String("journey", userId),
		zap.String("journey", "deleteUser"),
	)
	c.Status(http.StatusOK)
}
