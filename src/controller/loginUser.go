package controller

import (
	"net/http"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/validation"
	"github.com/Michael-Andryeer/crud-go/src/controller/model/request"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"github.com/Michael-Andryeer/crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser Controller",
		zap.String("journey", "LoginUser"),
		zap.String("operation", "init"),
	)

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "LoginUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)

	domainResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error(
			"Error trying to call LoginUser service",
			err,
			zap.String("journey", "LoginUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("LoginUser controller executed successfully",
		zap.String("journey", "LoginUser"),
		zap.String("journey", "LoginUser"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
