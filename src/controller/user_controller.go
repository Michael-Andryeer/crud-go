package controller

import (
	"github.com/Michael-Andryeer/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	DeleteUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}
