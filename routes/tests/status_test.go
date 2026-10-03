package routes_tests

import (
	"global-birthday-server/controllers"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"global-birthday-server/services"
	"global-birthday-server/startup"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusRoute(t *testing.T) {
	initialUsers := make([]models.User, 0)
	usersRepo := repositories.NewMockUsersController(&initialUsers)
	sessionsRepo := repositories.NewMockSessionsController(&initialUsers)
	sessionsService := services.NewSessionsService(sessionsRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)
	usersController := controllers.NewUsersController(usersService, sessionsService)
	statusController := controllers.NewStatusController()

	router := startup.SetupRouter(statusController, usersController)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/status/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Server is up.", w.Body.String())
}
