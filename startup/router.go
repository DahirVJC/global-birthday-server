package startup

import (
	"global-birthday-server/controllers"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"global-birthday-server/routes"
	"global-birthday-server/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	statusController := controllers.NewStatusController()

	initialUsers := make([]models.User, 0)
	usersRepo := repositories.NewMockUsersController(initialUsers)
	sessionsRepo := repositories.NewMockSessionsController(initialUsers)
	sessionsService := services.NewSessionsService(sessionsRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)
	usersController := controllers.NewUsersController(usersService, sessionsService)

	routes.RegisterRoutes(
		router,
		statusController,
		usersController,
	)

	return router
}

func StartRouter(router *gin.Engine) {
	err := router.Run("localhost:8080")

	if err != nil {
		return
	}
}
