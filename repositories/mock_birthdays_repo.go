package repositories

import (
	"fmt"
	"global-birthday-server/models"
	"slices"

	"github.com/google/uuid"
)

type MockBirthdaysRepository struct {
	birthdays *[]models.Birthday
}

func NewMockBirthdaysRepository(initialBirthdays *[]models.Birthday) MockBirthdaysRepository {
	return MockBirthdaysRepository{
		birthdays: initialBirthdays,
	}
}

func (br MockBirthdaysRepository) Get() (*[]models.Birthday, error) {
	return br.birthdays, nil
}

func (br MockBirthdaysRepository) GetById(id uuid.UUID) (*models.Birthday, error) {
	idx := slices.IndexFunc(*br.birthdays, func(b models.Birthday) bool { return b.ID == id })

	if idx == -1 {
		return nil, fmt.Errorf("birthday with ID %s not found", id)
	}

	return &(*br.birthdays)[idx], nil
}

func (br MockBirthdaysRepository) Create(birthday models.Birthday) error {
	newBirthdays := append(*br.birthdays, birthday)
	*br.birthdays = newBirthdays

	return nil
}

func (br MockBirthdaysRepository) Update(id uuid.UUID, birthday models.Birthday) error {
	idx := slices.IndexFunc(*br.birthdays, func(b models.Birthday) bool { return b.ID == id })

	if idx == -1 {
		return fmt.Errorf("birthday with ID %s not found", id)
	}

	(*br.birthdays)[idx] = birthday

	return nil
}

func (br MockBirthdaysRepository) Delete(id uuid.UUID) error {
	idx := slices.IndexFunc(*br.birthdays, func(b models.Birthday) bool { return b.ID == id })

	if idx == -1 {
		return fmt.Errorf("birthday with ID %s not found", id)
	}

	newBirthdays := slices.Delete(*br.birthdays, idx, idx+1)
	*br.birthdays = newBirthdays

	return nil
}

func (br MockBirthdaysRepository) Exists(id uuid.UUID) (bool, error) {
	return slices.ContainsFunc(*br.birthdays, func(birthday models.Birthday) bool {
		return birthday.ID == id
	}), nil
}
