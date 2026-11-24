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
	initialBirthdays := make([]models.Birthday, 0)
	initialEvents := make([]models.Event, 0)

	usersRepo := repositories.NewMockUsersRepository(&initialUsers)
	sessionsRepo := repositories.NewMockSessionsRepository(&initialUsers)
	birthdaysRepo := repositories.NewMockBirthdaysRepository(&initialBirthdays)
	eventsRepo := repositories.NewMockEventsRepository(&initialEvents)
	
	sessionsService := services.NewSessionsService(sessionsRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)
	birthdaysService := services.NewBirthdaysService(birthdaysRepo, usersRepo)
	eventsService := services.NewEventsService(eventsRepo, birthdaysRepo, usersRepo)
	
	usersController := controllers.NewUsersController(usersService, sessionsService)
	birthdaysController := controllers.NewBirthdaysController(birthdaysService, sessionsService)
	eventsController := controllers.NewEventsController(eventsService, sessionsService)

	router := startup.SetupRouter(statusController, usersController, birthdaysController, eventsController)
	startup.StartRouter(router)
}
