package routes_tests

import (
	"bytes"
	"encoding/json"
	"global-birthday-server/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBirthdaysRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/birthday/2026", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `[
		{
			"_id": "44444444-4444-4444-4444-444444444444",
			"name": "Kaia Ngata",
			"birthdate": "2000-01-15T00:00:00Z",
			"timezone": "Pacific/Auckland",
			"contactLinks": ["#kaina_555", "kaina@"],
			"wishlists": ["https://example.com/wishlist/kaia", "https://steamapps.com/wishlist/kaia"],
			"events": [],
			"startDate": "2026-01-14T06:00:00-05:00",
			"endDate": "2026-01-15T06:00:00-05:00"
		},
		{
			"_id": "55555555-5555-5555-5555-555555555555",
			"name": "Leap Day Friend",
			"birthdate": "2000-02-29T00:00:00Z",
			"timezone": "Pacific/Auckland",
			"contactLinks": [],
			"wishlists": [],
			"events": [],
			"startDate": "2026-02-27T06:00:00-05:00",
			"endDate": "2026-02-28T06:00:00-05:00"
		}
	]`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestGetBirthdayRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/birthday/2026/44444444-4444-4444-4444-444444444444", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `{
		"_id": "44444444-4444-4444-4444-444444444444",
		"name": "Kaia Ngata",
		"birthdate": "2000-01-15T00:00:00Z",
		"timezone": "Pacific/Auckland",
		"contactLinks": ["#kaina_555", "kaina@"],
		"wishlists": ["https://example.com/wishlist/kaia", "https://steamapps.com/wishlist/kaia"],
		"events": [],
		"startDate": "2026-01-14T06:00:00-05:00",
		"endDate": "2026-01-15T06:00:00-05:00"
	}`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestGetBirthdayLeapYearFallbackRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/birthday/2025/55555555-5555-5555-5555-555555555555", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `{
		"_id": "55555555-5555-5555-5555-555555555555",
		"name": "Leap Day Friend",
		"birthdate": "2000-02-29T00:00:00Z",
		"timezone": "Pacific/Auckland",
		"contactLinks": [],
		"wishlists": [],
		"events": [],
		"startDate": "2025-02-27T06:00:00-05:00",
		"endDate": "2025-02-28T06:00:00-05:00"
	}`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestGetBirthdayNotFoundRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/birthday/2026/77777777-7777-7777-7777-777777777777", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code, "Status code should be not found (GET)")
}

func TestGetBirthdayForbiddenAsNotFoundRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/birthday/2026/66666666-6666-6666-6666-666666666666", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code, "Status code should be not found (GET)")
}

func TestPostBirthdayRoute(t *testing.T) {
	initialUsers := getMockUsersWithSessions()
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	newBirthday := models.BirthdayReq{
		Name:      "Kaia Ngata",
		Birthdate: time.Date(1994, time.January, 15, 0, 0, 0, 0, time.UTC),
		Timezone:  "Pacific/Auckland",
		Wishlists: []string{"https://example.com/wishlist/kaia"},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(newBirthday)
	if err != nil {
		log.Fatal(err)
	}

	postRecorder := httptest.NewRecorder()
	postReq, _ := http.NewRequest(http.MethodPost, "/api/birthday/", &buf)
	postReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(postRecorder, postReq)

	require.Equal(t, http.StatusCreated, postRecorder.Code, "Status code should be Created (POST)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/birthday/2026", nil)
	getReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusOK, getRecorder.Code, "Status code should be OK (GET)")

	var birthdays []models.BirthdayRes
	if err := json.NewDecoder(getRecorder.Body).Decode(&birthdays); err != nil {
		log.Fatal(err)
	}

	require.Equal(t, 1, len(birthdays), "Should persist one birthday")
	assert.Equal(t, "Kaia Ngata", birthdays[0].Name)
	assert.Equal(t, "Pacific/Auckland", birthdays[0].Timezone)
	assert.Equal(t, 2000, birthdays[0].Birthdate.Year())
	assert.Equal(t, time.January, birthdays[0].Birthdate.Month())
	assert.Equal(t, 15, birthdays[0].Birthdate.Day())
	assert.Equal(t, []string{"https://example.com/wishlist/kaia"}, birthdays[0].Wishlists)
	assert.Equal(t, "2026-01-14T06:00:00-05:00", birthdays[0].StartDate.Format(time.RFC3339))
	assert.Equal(t, "2026-01-15T06:00:00-05:00", birthdays[0].EndDate.Format(time.RFC3339))
}

func TestPutBirthdayRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	updatedBirthday := models.BirthdayReq{
		Name:      "Kaia Updated",
		Birthdate: time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC),
		Timezone:  "Pacific/Auckland",
		Wishlists: []string{"https://example.com/wishlist/updated"},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedBirthday)
	if err != nil {
		log.Fatal(err)
	}

	putRecorder := httptest.NewRecorder()
	putReq, _ := http.NewRequest(http.MethodPut, "/api/birthday/44444444-4444-4444-4444-444444444444", &buf)
	putReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(putRecorder, putReq)

	require.Equal(t, http.StatusOK, putRecorder.Code, "Status code should be OK (PUT)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/birthday/2026/44444444-4444-4444-4444-444444444444", nil)
	getReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusOK, getRecorder.Code, "Status code should be OK (GET)")

	expectedJSON := `{
		"_id": "44444444-4444-4444-4444-444444444444",
		"name": "Kaia Updated",
		"birthdate": "2000-01-15T00:00:00Z",
		"timezone": "Pacific/Auckland",
		"contactLinks": [],
		"wishlists": ["https://example.com/wishlist/updated"],
		"events": [],
		"startDate": "2026-01-14T06:00:00-05:00",
		"endDate": "2026-01-15T06:00:00-05:00"
	}`

	assert.JSONEq(t, expectedJSON, getRecorder.Body.String())
}

func TestDeleteBirthdayRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	deleteRecorder := httptest.NewRecorder()
	deleteReq, _ := http.NewRequest(http.MethodDelete, "/api/birthday/44444444-4444-4444-4444-444444444444", nil)
	deleteReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(deleteRecorder, deleteReq)

	require.Equal(t, http.StatusOK, deleteRecorder.Code, "Status code should be OK (DELETE)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/birthday/2026/44444444-4444-4444-4444-444444444444", nil)
	getReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(getRecorder, getReq)

	assert.Equal(t, http.StatusNotFound, getRecorder.Code, "Status code should be 404 (GET)")
}
