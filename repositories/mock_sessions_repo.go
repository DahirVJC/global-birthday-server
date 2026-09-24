package repositories

import (
	"fmt"
	"global-birthday-server/models"
	"slices"

	"github.com/google/uuid"
)

type MockSessionsRepository struct {
	users []models.User
}

func NewMockSessionsController(initialUsers []models.User) MockSessionsRepository {
	return MockSessionsRepository{
		users: initialUsers,
	}
}

func (sr MockSessionsRepository) Validate(hashedToken string) bool {
	return slices.ContainsFunc(sr.users, func(user models.User) bool {
		return slices.ContainsFunc(user.Sessions, func(session models.Session) bool {
			return session.AccessToken == hashedToken
		})
	})
}

func (sr MockSessionsRepository) GetUserId(hashedToken string) (uuid.UUID, error) {
	idx := slices.IndexFunc(sr.users, func(user models.User) bool {
		return slices.ContainsFunc(user.Sessions, func(session models.Session) bool {
			return session.AccessToken == hashedToken
		})
	})

	if idx == -1 {
		return uuid.Nil, fmt.Errorf(
			"user with session ID %s not found",
			hashedToken,
		)
	}

	return sr.users[idx].ID, nil
}

func (sr MockSessionsRepository) Add(userId uuid.UUID, session models.Session) error {
	idx := slices.IndexFunc(sr.users, func(user models.User) bool {
		return user.ID == userId
	})

	if idx == -1 {
		return fmt.Errorf(
			"user with ID %s not found",
			userId,
		)
	}

	sr.users[idx].Sessions = append(sr.users[idx].Sessions, session)

	return nil
}
