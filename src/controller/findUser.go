package controller

import (
	"net/http"
	"net/mail"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller",
		zap.String("journey", "FindUserById"),
	)
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying	 to validate user id", err,
			zap.String("journey", "FindUserById"),
		)

		errorMessage := rest_errors.NewBadRequestError(
			"User ID is not a valid ID",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call FindUserByIDServices service", err,
			zap.String("journey", "FindUserById"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "FindUserById"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller",
		zap.String("journey", "FindUserByEmail"),
	)

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate user email", err,
			zap.String("journey", "FindUserByEmail"),
		)
		errorMessage := rest_errors.NewBadRequestError(
			"User email is not valid",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmailServices service", err,
			zap.String("journey", "FindUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("userEmail", userEmail),
		zap.String("journey", "FindUserByEmail"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
