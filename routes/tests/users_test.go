package routes_tests

import (
	"bytes"
	"encoding/json"
	"global-birthday-server/controllers"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"global-birthday-server/services"
	"global-birthday-server/startup"
	"log"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
)

func getMockUsers() []models.User {
	return []models.User{
		{
			ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:      "Alice Nguyen",
			Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
			Email:     "alice.nguyen@example.com",
			Timezone:  "America/New_York",
			Birthdays: nil,
			Sessions:  nil,
		},
		{
			ID:        uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:      "Bob Mercado",
			Password:  "$2a$10$X7trdoO6c2Ym0XWR1vNXVeWbnqZ6yNlTeqrMc0BpeDBMdCvQRtszi",
			Email:     "bob.mercado@example.com",
			Timezone:  "Europe/Lisbon",
			Birthdays: nil,
			Sessions:  nil,
		},
		{
			ID:        uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			Name:      "Chiemi Tanaka",
			Password:  "$2a$10$cZ4gVKFmqkhRRRUnA5MTYuQcwBTmEQT3wjQMxjOkCEwp0GT2EfZ1O",
			Email:     "chiemi.tanaka@example.co.jp",
			Timezone:  "Asia/Tokyo",
			Birthdays: nil,
			Sessions:  nil,
		},
	}
}

func TestGetUserRoute(t *testing.T) {
	initialUsers := getMockUsers()
	usersRepo := repositories.NewMockUsersController(&initialUsers)
	sessionsRepo := repositories.NewMockSessionsController(&initialUsers)
	sessionsService := services.NewSessionsService(sessionsRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)
	usersController := controllers.NewUsersController(usersService, sessionsService)
	statusController := controllers.NewStatusController()

	router := startup.SetupRouter(statusController, usersController)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/users/", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `[
        {
            "_id": "11111111-1111-1111-1111-111111111111",
            "name": "Alice Nguyen",
            "email": "alice.nguyen@example.com",
            "timezone": "America/New_York"
        },
        {
            "_id": "22222222-2222-2222-2222-222222222222",
            "name": "Bob Mercado",
            "email": "bob.mercado@example.com",
            "timezone": "Europe/Lisbon"
        },
        {
            "_id": "33333333-3333-3333-3333-333333333333",
            "name": "Chiemi Tanaka",
            "email": "chiemi.tanaka@example.co.jp",
            "timezone": "Asia/Tokyo"
        }
    ]`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}
func TestPostUserRoute(t *testing.T) {
	initialUsers := getMockUsers()
	usersRepo := repositories.NewMockUsersController(&initialUsers)
	sessionsRepo := repositories.NewMockSessionsController(&initialUsers)
	sessionsService := services.NewSessionsService(sessionsRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)
	usersController := controllers.NewUsersController(usersService, sessionsService)
	statusController := controllers.NewStatusController()

	router := startup.SetupRouter(statusController, usersController)

	newUser := models.UserReq{
		Name:     "User Test",
		Password: "!15b2DwMEC^o9D",
		Email:    "test@example.com",
		Timezone: "America/Los_Angeles",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(newUser)
	if err != nil {
		log.Fatal(err)
	}

	postRecorder := httptest.NewRecorder()
	postReq, _ := http.NewRequest(http.MethodPost, "/api/users/", &buf)
	router.ServeHTTP(postRecorder, postReq)

	require.Equal(t, http.StatusOK, postRecorder.Code, "Status code should be OK (POST)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/users/", nil)
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusOK, getRecorder.Code, "Status code should be OK (GET)")

	var users []models.UserRes

	if err := json.NewDecoder(getRecorder.Body).Decode(&users); err != nil {
		log.Fatal(err)
	}

	expectedUsers := []models.UserReq{
		{
			Name:     "Alice Nguyen",
			Email:    "alice.nguyen@example.com",
			Timezone: "America/New_York",
		},
		{
			Name:     "Bob Mercado",
			Email:    "bob.mercado@example.com",
			Timezone: "Europe/Lisbon",
		},
		{
			Name:     "Chiemi Tanaka",
			Email:    "chiemi.tanaka@example.co.jp",
			Timezone: "Asia/Tokyo",
		},
		{
			Name:     "User Test",
			Email:    "test@example.com",
			Timezone: "America/Los_Angeles",
		},
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})

	sort.Slice(expectedUsers, func(i, j int) bool {
		return expectedUsers[i].Name < expectedUsers[j].Name
	})

	require.Equal(t, len(expectedUsers), len(users), "Res Users and Expected Users should have the same length")

	for idx, user := range users {
		assert.Equal(t, expectedUsers[idx].Name, user.Name, "User name should be equal")
		assert.Equal(t, expectedUsers[idx].Email, user.Email, "User email should be equal")
		assert.Equal(t, expectedUsers[idx].Timezone, user.Timezone, "User timezone should be equal")
	}
}
