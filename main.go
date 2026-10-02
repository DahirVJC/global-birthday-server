package main

import (
	"global-birthday-server/controllers"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"global-birthday-server/services"
	"global-birthday-server/startup"
)

func main() {
	startup.LoadEnvironment()

	statusController := controllers.NewStatusController()

	initialUsers := make([]models.User, 0)
	usersRepo := repositories.NewMockUsersController(initialUsers)
	sessionsRepo := repositories.NewMockSessionsController(initialUsers)
	sessionsService := services.NewSessionsService(sessionsRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)
	usersController := controllers.NewUsersController(usersService, sessionsService)

	router := startup.SetupRouter(statusController, usersController)
	startup.StartRouter(router)
}
