package controller

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := rest_errors.NewInternalServerError("Error creating user")
	c.JSON(err.Code, err)
}
