package controller

import (
	"fmt"

	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_errors.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, error=%s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
