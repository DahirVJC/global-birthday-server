package services

import (
	"global-birthday-server/models"
	"global-birthday-server/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUsersService() *UserService {
	mongoRepo := repositories.NewUsersController()
	return &UserService{
		repo: mongoRepo,
	}
}

func (u UserService) GetById(id uuid.UUID) models.UserRes {
	var user models.UserRes
	return user
}

func (u UserService) Get() []models.UserRes {
	var users []models.UserRes
	return users
}

func (u UserService) Create(user models.UserReq) (models.UserReq, error) {
	return user, nil
}

func (u UserService) Update(id uuid.UUID, user models.UserReq) error {
	return nil
}

func (u UserService) Delete(id uuid.UUID) error {
	return nil
}
