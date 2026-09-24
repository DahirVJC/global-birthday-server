package repositories

import (
	"fmt"
	"global-birthday-server/models"
	"slices"

	"github.com/google/uuid"
)

type MockUsersRepository struct {
	users []models.User
}

func NewMockUsersController(initialUsers []models.User) MockUsersRepository {
	return MockUsersRepository{
		users: initialUsers,
	}
}

func toUserRef(users []models.User) []*models.User {
	result := make([]*models.User, 0, len(users))

	for i := range users {
		result = append(result, &users[i])
	}

	return result
}

func (ur MockUsersRepository) Get() ([]*models.User, error) {
	refUsers := toUserRef(ur.users)

	return refUsers, nil
}

func (ur MockUsersRepository) GetById(id uuid.UUID) (*models.User, error) {
	refUsers := toUserRef(ur.users)

	idx := slices.IndexFunc(refUsers, func(u *models.User) bool { return u.ID == id })

	if idx == -1 {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}

	return refUsers[idx], nil
}

func (ur MockUsersRepository) Create(user models.User) error {
	ur.users = append(ur.users, user)

	return nil
}

func (ur MockUsersRepository) Update(id uuid.UUID, user models.User) error {
	idx := slices.IndexFunc(ur.users, func(u models.User) bool { return u.ID == id })

	if idx == -1 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	ur.users[idx] = user

	return nil
}

func (ur MockUsersRepository) Delete(id uuid.UUID) error {
	idx := slices.IndexFunc(ur.users, func(u models.User) bool { return u.ID == id })

	if idx == -1 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	ur.users = slices.Delete(ur.users, idx, idx+1)

	return nil
}

func (ur MockUsersRepository) Exists(id uuid.UUID) (bool, error) {
	return slices.ContainsFunc(ur.users, func(user models.User) bool {
		return user.ID == id
	}), nil
}
