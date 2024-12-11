package routes

import (
	"github.com/Michael-Andryeer/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

// inicializar as rotas da aplicação
func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUserById)
}
