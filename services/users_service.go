package services

import (
	"global-birthday-server/models"

	"github.com/google/uuid"
)

type UserService struct{}

func (u UserService) GetById(id uuid.UUID) models.User {
	var user models.User
	return user
}

func (u UserService) Get() []models.User {
	var users []models.User
	return users
}

func (u UserService) Create(user models.User) (models.User, error) {
	return user, nil
}

func (u UserService) Update(id uuid.UUID) error {
	return nil
}

func (u UserService) Delete(id uuid.UUID) error {
	return nil
}
