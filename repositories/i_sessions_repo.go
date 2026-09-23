package repositories

import (
	"global-birthday-server/models"

	"github.com/google/uuid"
)

type SessionsRepository interface {
	Validate(accessToken string) bool
	GetUserId(accessToken string) (uuid.UUID, error)
	Add(userId uuid.UUID, session models.Session) error
}
