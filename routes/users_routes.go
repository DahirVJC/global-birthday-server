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
	}
}
