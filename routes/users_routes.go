package routes

import (
	"global-birthday-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(
	router *gin.RouterGroup,
	usersController *controllers.UsersController) {
	users := router.Group("/users")
	{
		users.GET("/", usersController.GetUsers)
		users.GET("/me", usersController.GetUser)
		users.POST("/", usersController.PostUser)
		users.PUT("/me", usersController.PutUser)
		users.DELETE("/me", usersController.DeleteUser)
	}
}
