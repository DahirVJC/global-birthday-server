package repositories

import (
	"fmt"
	"global-birthday-server/models"
	"slices"

	"github.com/google/uuid"
)

type MockUsersRepository struct {
	users *[]models.User
}

func NewMockUsersRepository(initialUsers *[]models.User) MockUsersRepository {
	return MockUsersRepository{
		users: initialUsers,
	}
}

func (ur MockUsersRepository) Get() (*[]models.User, error) {
	return ur.users, nil
}

func (ur MockUsersRepository) GetById(id uuid.UUID) (*models.User, error) {
	idx := slices.IndexFunc(*ur.users, func(u models.User) bool { return u.ID == id })

	if idx == -1 {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}

	return &(*ur.users)[idx], nil
}

func (ur MockUsersRepository) Create(user models.User) error {
	newUsers := append(*ur.users, user)
	*ur.users = newUsers

	return nil
}

func (ur MockUsersRepository) Update(id uuid.UUID, user models.User) error {
	idx := slices.IndexFunc(*ur.users, func(u models.User) bool { return u.ID == id })

	if idx == -1 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	(*ur.users)[idx] = user

	return nil
}

func (ur MockUsersRepository) Delete(id uuid.UUID) error {
	idx := slices.IndexFunc(*ur.users, func(u models.User) bool { return u.ID == id })

	if idx == -1 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	newUsers := slices.Delete(*ur.users, idx, idx+1)
	*ur.users = newUsers

	return nil
}

func (ur MockUsersRepository) Exists(id uuid.UUID) (bool, error) {
	return slices.ContainsFunc(*ur.users, func(user models.User) bool {
		return user.ID == id
	}), nil
}
