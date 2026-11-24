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

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEventsRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdaysWithEvents()
	initialEvents := getMockEvents()
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/event/", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `[
		{
			"_id": "99999999-9999-9999-9999-999999999999",
			"birthdayId": "44444444-4444-4444-4444-444444444444",
			"eventDateStart": "2030-01-16T12:00:00Z",
			"eventDateEnd": "2030-01-16T14:00:00Z",
			"description": "Kaia party",
			"startDate": "2030-01-15T18:00:00-05:00",
			"endDate": "2030-01-15T20:00:00-05:00"
		}
	]`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestGetEventsIncludingPastRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdaysWithEvents()
	initialEvents := getMockEvents()
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/event/?past=true", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `[
		{
			"_id": "88888888-8888-8888-8888-888888888888",
			"birthdayId": "44444444-4444-4444-4444-444444444444",
			"eventDateStart": "2020-01-16T12:00:00Z",
			"eventDateEnd": "2020-01-16T14:00:00Z",
			"description": "Past party",
			"startDate": "2020-01-15T18:00:00-05:00",
			"endDate": "2020-01-15T20:00:00-05:00"
		},
		{
			"_id": "99999999-9999-9999-9999-999999999999",
			"birthdayId": "44444444-4444-4444-4444-444444444444",
			"eventDateStart": "2030-01-16T12:00:00Z",
			"eventDateEnd": "2030-01-16T14:00:00Z",
			"description": "Kaia party",
			"startDate": "2030-01-15T18:00:00-05:00",
			"endDate": "2030-01-15T20:00:00-05:00"
		}
	]`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestGetEventRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdaysWithEvents()
	initialEvents := getMockEvents()
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/event/99999999-9999-9999-9999-999999999999", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "Status code should be OK (GET)")

	expectedJSON := `{
		"_id": "99999999-9999-9999-9999-999999999999",
		"birthdayId": "44444444-4444-4444-4444-444444444444",
		"eventDateStart": "2030-01-16T12:00:00Z",
		"eventDateEnd": "2030-01-16T14:00:00Z",
		"description": "Kaia party",
		"startDate": "2030-01-15T18:00:00-05:00",
		"endDate": "2030-01-15T20:00:00-05:00"
	}`

	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestGetEventNotFoundRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdaysWithEvents()
	initialEvents := getMockEvents()
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/event/77777777-7777-7777-7777-777777777777", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code, "Status code should be not found (GET)")
}

func TestGetEventForbiddenAsNotFoundRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdaysWithEvents()
	initialEvents := getMockEvents()
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/event/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", nil)
	req.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code, "Status code should be not found (GET)")
}

func TestPostEventRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdays()
	initialEvents := make([]models.Event, 0)
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	end := time.Date(2030, time.January, 16, 14, 0, 0, 0, time.UTC)
	description := "Kaia party"
	newEvent := models.EventReq{
		BirthdayID:     uuid.MustParse("44444444-4444-4444-4444-444444444444"),
		EventDateStart: time.Date(2030, time.January, 16, 12, 0, 0, 0, time.UTC),
		EventDateEnd:   &end,
		Description:    &description,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(newEvent)
	if err != nil {
		log.Fatal(err)
	}

	postRecorder := httptest.NewRecorder()
	postReq, _ := http.NewRequest(http.MethodPost, "/api/event/", &buf)
	postReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(postRecorder, postReq)

	require.Equal(t, http.StatusCreated, postRecorder.Code, "Status code should be Created (POST)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/event/", nil)
	getReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusOK, getRecorder.Code, "Status code should be OK (GET)")

	var events []models.EventRes
	if err := json.NewDecoder(getRecorder.Body).Decode(&events); err != nil {
		log.Fatal(err)
	}

	require.Equal(t, 1, len(events), "Should persist one event")
	assert.Equal(t, uuid.MustParse("44444444-4444-4444-4444-444444444444"), events[0].BirthdayID)
	assert.Equal(t, "Kaia party", *events[0].Description)
	assert.Equal(t, "2030-01-15T18:00:00-05:00", events[0].StartDate.Format(time.RFC3339))
	assert.Equal(t, "2030-01-15T20:00:00-05:00", events[0].EndDate.Format(time.RFC3339))

	birthdayRecorder := httptest.NewRecorder()
	birthdayReq, _ := http.NewRequest(http.MethodGet, "/api/birthday/2030/44444444-4444-4444-4444-444444444444", nil)
	birthdayReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(birthdayRecorder, birthdayReq)

	require.Equal(t, http.StatusOK, birthdayRecorder.Code)
	var birthday models.BirthdayRes
	if err := json.NewDecoder(birthdayRecorder.Body).Decode(&birthday); err != nil {
		log.Fatal(err)
	}
	require.Equal(t, 1, len(birthday.Events), "Birthday should reference the new event")
	assert.Equal(t, events[0].ID, birthday.Events[0])
}

func TestPutEventRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdaysWithEvents()
	initialEvents := getMockEvents()
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	end := time.Date(2030, time.January, 16, 16, 0, 0, 0, time.UTC)
	description := "Updated party"
	updatedEvent := models.EventReq{
		BirthdayID:     uuid.MustParse("44444444-4444-4444-4444-444444444444"),
		EventDateStart: time.Date(2030, time.January, 16, 12, 0, 0, 0, time.UTC),
		EventDateEnd:   &end,
		Description:    &description,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedEvent)
	if err != nil {
		log.Fatal(err)
	}

	putRecorder := httptest.NewRecorder()
	putReq, _ := http.NewRequest(http.MethodPut, "/api/event/99999999-9999-9999-9999-999999999999", &buf)
	putReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(putRecorder, putReq)

	require.Equal(t, http.StatusOK, putRecorder.Code, "Status code should be OK (PUT)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/event/99999999-9999-9999-9999-999999999999", nil)
	getReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusOK, getRecorder.Code, "Status code should be OK (GET)")

	expectedJSON := `{
		"_id": "99999999-9999-9999-9999-999999999999",
		"birthdayId": "44444444-4444-4444-4444-444444444444",
		"eventDateStart": "2030-01-16T12:00:00Z",
		"eventDateEnd": "2030-01-16T16:00:00Z",
		"description": "Updated party",
		"startDate": "2030-01-15T18:00:00-05:00",
		"endDate": "2030-01-15T22:00:00-05:00"
	}`

	assert.JSONEq(t, expectedJSON, getRecorder.Body.String())
}

func TestDeleteEventRoute(t *testing.T) {
	initialUsers := getMockUsersWithBirthdays()
	initialBirthdays := getMockBirthdaysWithEvents()
	initialEvents := getMockEvents()
	router := setupTestRouterWithEvents(&initialUsers, &initialBirthdays, &initialEvents)

	deleteRecorder := httptest.NewRecorder()
	deleteReq, _ := http.NewRequest(http.MethodDelete, "/api/event/99999999-9999-9999-9999-999999999999", nil)
	deleteReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(deleteRecorder, deleteReq)

	require.Equal(t, http.StatusOK, deleteRecorder.Code, "Status code should be OK (DELETE)")

	getRecorder := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, "/api/event/99999999-9999-9999-9999-999999999999", nil)
	getReq.Header.Set("access-token", aliceAccessToken)
	router.ServeHTTP(getRecorder, getReq)

	assert.Equal(t, http.StatusNotFound, getRecorder.Code, "Status code should be 404 (GET)")
}
