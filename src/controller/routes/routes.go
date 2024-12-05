package routes

import (
	"github.com/Michael-Andryeer/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

// inicializar as rotas da aplicação
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUserById)
}
