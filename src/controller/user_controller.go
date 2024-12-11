package controller

import (
	"github.com/Michael-Andryeer/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{}
}

type UserController interface {
	DeleteUser(c *gin.Context)
	FindUser(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
