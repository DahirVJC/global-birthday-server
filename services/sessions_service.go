package services

import (
	"fmt"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"global-birthday-server/utilities"
	"time"

	"github.com/google/uuid"
)

type SessionsService struct {
	sessionsRepo repositories.SessionsRepository
	usersRepo    repositories.UsersRepository
}

func NewSessionsService(sessionsRepo repositories.SessionsRepository, usersRepo repositories.UsersRepository) *SessionsService {
	return &SessionsService{
		sessionsRepo: sessionsRepo,
		usersRepo:    usersRepo,
	}
}

func (ss SessionsService) Validate(accessToken string) bool {
	return ss.sessionsRepo.Validate(utilities.ToHash(accessToken))
}

func (ss SessionsService) GetUserId(accessToken string) (uuid.UUID, error) {
	return ss.sessionsRepo.GetUserId(utilities.ToHash(accessToken))
}

func (ss SessionsService) Create(userId uuid.UUID) (uuid.UUID, error) {
	userExists, err := ss.usersRepo.Exists(userId)

	if err != nil {
		return uuid.Nil, err
	}

	if !userExists {
		return uuid.Nil, fmt.Errorf("user with ID %s not found", userId)
	}

	token, err := utilities.GenerateToken()
	if err != nil {
		return uuid.Nil, err
	}

	newSession := models.Session{
		ID:          uuid.New(),
		AccessToken: utilities.ToHash(token), // Store only the hash
		CreatedAt:   time.Now().UTC(),
		ExpiresAt:   time.Now().UTC().Add(90 * 24 * time.Hour),
	}

	if err := ss.sessionsRepo.Add(userId, newSession); err != nil {
		return uuid.Nil, err
	}

	return newSession.ID, nil
}
