package routes_tests

import (
	"global-birthday-server/controllers"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"global-birthday-server/services"
	"global-birthday-server/startup"

	"github.com/gin-gonic/gin"
)

func setupTestRouter(users *[]models.User, birthdays *[]models.Birthday) *gin.Engine {
	usersRepo := repositories.NewMockUsersRepository(users)
	sessionsRepo := repositories.NewMockSessionsRepository(users)
	birthdaysRepo := repositories.NewMockBirthdaysRepository(birthdays)
	sessionsService := services.NewSessionsService(sessionsRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)
	birthdaysService := services.NewBirthdaysService(birthdaysRepo, usersRepo)
	usersController := controllers.NewUsersController(usersService, sessionsService)
	birthdaysController := controllers.NewBirthdaysController(birthdaysService, sessionsService)
	statusController := controllers.NewStatusController()

	return startup.SetupRouter(statusController, usersController, birthdaysController)
}
