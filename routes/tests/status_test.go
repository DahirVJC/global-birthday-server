package routes_tests

import (
	"global-birthday-server/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusRoute(t *testing.T) {
	initialUsers := make([]models.User, 0)
	initialBirthdays := make([]models.Birthday, 0)
	router := setupTestRouter(&initialUsers, &initialBirthdays)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/status/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Server is up.", w.Body.String())
}
