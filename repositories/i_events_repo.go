package repositories

import (
	"global-birthday-server/models"

	"github.com/google/uuid"
)

type EventsRepository interface {
	Get() (*[]models.Event, error)
	GetById(id uuid.UUID) (*models.Event, error)
	Create(event models.Event) error
	Update(id uuid.UUID, event models.Event) error
	Delete(id uuid.UUID) error
	Exists(id uuid.UUID) (bool, error)
}
