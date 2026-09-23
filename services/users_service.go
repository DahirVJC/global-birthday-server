package services

import (
	"global-birthday-server/models"
	"global-birthday-server/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUsersService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func toUserRes(userDb *models.User) *models.UserRes {
	if userDb == nil {
		return nil
	}

	res := &models.UserRes{
		ID:       userDb.ID,
		Name:     userDb.Name,
		Email:    userDb.Email,
		Timezone: userDb.Timezone,
	}

	return res
}

func toUserDb(userReq models.UserReq) models.User {
	user := models.User{
		ID:        uuid.New(),
		Name:      userReq.Name,
		Password:  userReq.Password,
		Email:     userReq.Email,
		Timezone:  userReq.Timezone,
		Birthdays: nil,
		Sessions:  nil,
	}

	return user
}

func (us UserService) GetById(id uuid.UUID) (*models.UserRes, error) {
	userDb, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return toUserRes(userDb), nil
}

func (us UserService) Get() ([]*models.UserRes, error) {
	usersDb, err := us.repo.Get()

	if err != nil {
		return nil, err
	}

	var users []*models.UserRes

	for _, userDb := range usersDb {
		users = append(users, toUserRes(userDb))
	}

	return users, nil
}

func (us UserService) Create(user models.UserReq) error {
	userDb := toUserDb(user)

	return us.repo.Create(userDb)
}

func (us UserService) Update(id uuid.UUID, user models.UserReq) error {
	userDb := toUserDb(user)

	return us.repo.Update(id, userDb)
}

func (us UserService) Delete(id uuid.UUID) error {
	return us.repo.Delete(id)
}
