package repositories

import (
	"fmt"
	"global-birthday-server/models"
	"slices"

	"github.com/google/uuid"
)

type MockEventsRepository struct {
	events *[]models.Event
}

func NewMockEventsRepository(initialEvents *[]models.Event) MockEventsRepository {
	return MockEventsRepository{
		events: initialEvents,
	}
}

func (er MockEventsRepository) Get() (*[]models.Event, error) {
	return er.events, nil
}

func (er MockEventsRepository) GetById(id uuid.UUID) (*models.Event, error) {
	idx := slices.IndexFunc(*er.events, func(e models.Event) bool { return e.ID == id })

	if idx == -1 {
		return nil, fmt.Errorf("event with ID %s not found", id)
	}

	return &(*er.events)[idx], nil
}

func (er MockEventsRepository) Create(event models.Event) error {
	newEvents := append(*er.events, event)
	*er.events = newEvents

	return nil
}

func (er MockEventsRepository) Update(id uuid.UUID, event models.Event) error {
	idx := slices.IndexFunc(*er.events, func(e models.Event) bool { return e.ID == id })

	if idx == -1 {
		return fmt.Errorf("event with ID %s not found", id)
	}

	(*er.events)[idx] = event

	return nil
}

func (er MockEventsRepository) Delete(id uuid.UUID) error {
	idx := slices.IndexFunc(*er.events, func(e models.Event) bool { return e.ID == id })

	if idx == -1 {
		return fmt.Errorf("event with ID %s not found", id)
	}

	newEvents := slices.Delete(*er.events, idx, idx+1)
	*er.events = newEvents

	return nil
}

func (er MockEventsRepository) Exists(id uuid.UUID) (bool, error) {
	return slices.ContainsFunc(*er.events, func(event models.Event) bool {
		return event.ID == id
	}), nil
}
