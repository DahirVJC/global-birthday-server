package routes

import (
	"global-birthday-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	statusController *controllers.StatusController,
	usersController *controllers.UsersController) {
	api := router.Group("/api")

	RegisterStatusRoutes(api, statusController)
	RegisterUsersRoutes(api, usersController)
}
