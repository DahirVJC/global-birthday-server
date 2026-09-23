package repositories

import (
	"global-birthday-server/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Get() ([]*models.User, error)
	GetById(id uuid.UUID) (*models.User, error)
	Create(user models.User) error
	Update(id uuid.UUID, user models.User) error
	Delete(id uuid.UUID) error
}
