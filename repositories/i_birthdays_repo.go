package repositories

import (
	"global-birthday-server/models"

	"github.com/google/uuid"
)

type BirthdaysRepository interface {
	Get() (*[]models.Birthday, error)
	GetById(id uuid.UUID) (*models.Birthday, error)
	Create(birthday models.Birthday) error
	Update(id uuid.UUID, birthday models.Birthday) error
	Delete(id uuid.UUID) error
	Exists(id uuid.UUID) (bool, error)
}
