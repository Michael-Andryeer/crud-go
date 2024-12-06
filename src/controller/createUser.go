package controller

import (
	"net/http"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/validation"
	"github.com/Michael-Andryeer/crud-go/src/controller/model/request"
	"github.com/Michael-Andryeer/crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	)
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zapcore.Field{
				Key:    "journey",
				String: "createUser",
			},
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	response := response.UserResponse{
		Id:    "Test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	)
	c.JSON(http.StatusOK, response)
}
