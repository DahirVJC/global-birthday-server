package routes

import (
	"global-birthday-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(
	router *gin.RouterGroup,
	userController *controllers.UsersController) {
	users := router.Group("/users")
	{
		users.GET("/", userController.GetUsers)
		users.GET("/me", userController.GetUser)
		users.POST("/", userController.PostUser)
		users.PUT("/me")
		users.DELETE("/me", userController.DeleteUser)
	}
}
