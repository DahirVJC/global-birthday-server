package routes_tests

import (
	"bytes"
	"encoding/json"
	"global-birthday-server/models"
	"log"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsersRoute(t *testing.T) {
	initialUsers := getMockUsers()
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)

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

func TestGetUserRoute(t *testing.T) {
	initialUsers := getMockUsersWithSessions()
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/users/me", nil)
	req.Header.Set("access-token", "8e1jgIYXjcoHvNDFmvKff0RPQBFGtZZl4nqK4u_OkRI")
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `
        {
            "_id": "11111111-1111-1111-1111-111111111111",
            "name": "Alice Nguyen",
            "email": "alice.nguyen@example.com",
            "timezone": "America/New_York"
        }`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestGetUserNotFoundRoute(t *testing.T) {
	initialUsers := getMockUsersWithSessions()
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/users/me", nil)
	req.Header.Set("access-token", "123")
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code, "Status code should be not found (GET)")
}

func TestPostUserRoute(t *testing.T) {
	initialUsers := getMockUsers()
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)

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

func TestPutUserRoute(t *testing.T) {
	initialUsers := getMockUsersWithSessions()
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)

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

	putRecorder := httptest.NewRecorder()
	putReq, _ := http.NewRequest(http.MethodPut, "/api/users/me", &buf)
	putReq.Header.Set("access-token", "8e1jgIYXjcoHvNDFmvKff0RPQBFGtZZl4nqK4u_OkRI")
	router.ServeHTTP(putRecorder, putReq)

	require.Equal(t, http.StatusOK, putRecorder.Code, "Status code should be OK (PUT)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/users/me", nil)
	getReq.Header.Set("access-token", "8e1jgIYXjcoHvNDFmvKff0RPQBFGtZZl4nqK4u_OkRI")
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusOK, getRecorder.Code, "Status code should be OK (GET)")

	expectedJSON := `
        {
            "_id": "11111111-1111-1111-1111-111111111111",
            "name": "User Test",
            "email": "test@example.com",
            "timezone": "America/Los_Angeles"
        }`

	assert.JSONEq(t, expectedJSON, getRecorder.Body.String())
}

func TestDeleteUserRoute(t *testing.T) {
	initialUsers := getMockUsersWithSessions()
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	deleteRecorder := httptest.NewRecorder()
	deleteReq, _ := http.NewRequest(http.MethodDelete, "/api/users/me", nil)
	deleteReq.Header.Set("access-token", "8e1jgIYXjcoHvNDFmvKff0RPQBFGtZZl4nqK4u_OkRI")
	router.ServeHTTP(deleteRecorder, deleteReq)

	require.Equal(t, http.StatusOK, deleteRecorder.Code, "Status code should be OK (DELETE)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/users/me", nil)
	getReq.Header.Set("access-token", "8e1jgIYXjcoHvNDFmvKff0RPQBFGtZZl4nqK4u_OkRI")
	router.ServeHTTP(getRecorder, getReq)

	assert.Equal(t, http.StatusNotFound, getRecorder.Code, "Status code should be 404 (GET)")
}
