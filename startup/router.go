package startup

import (
	"global-birthday-server/controllers"
	"global-birthday-server/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	statusController *controllers.StatusController,
	usersController *controllers.UsersController,
	birthdaysController *controllers.BirthdaysController,
	eventsController *controllers.EventsController) *gin.Engine {
	router := gin.Default()

	routes.RegisterRoutes(
		router,
		statusController,
		usersController,
		birthdaysController,
		eventsController,
	)

	return router
}

func StartRouter(router *gin.Engine) {
	err := router.Run("localhost:8080")

	if err != nil {
		return
	}
}
