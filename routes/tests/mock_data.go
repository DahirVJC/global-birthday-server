package routes_tests

import (
	"global-birthday-server/models"
	"global-birthday-server/utilities"
	"time"

	"github.com/google/uuid"
)

const aliceAccessToken = "8e1jgIYXjcoHvNDFmvKff0RPQBFGtZZl4nqK4u_OkRI"

func getMockBirthdays() []models.Birthday {
	return []models.Birthday{
		{
			ID:           uuid.MustParse("44444444-4444-4444-4444-444444444444"),
			Name:         "Kaia Ngata",
			Birthdate:    time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC),
			Timezone:     "Pacific/Auckland",
			ContactLinks: []string{"#kaina_555", "kaina@"},
			Wishlists:    []string{"https://example.com/wishlist/kaia", "https://steamapps.com/wishlist/kaia"},
			Events:       []uuid.UUID{},
		},
		{
			ID:           uuid.MustParse("55555555-5555-5555-5555-555555555555"),
			Name:         "Leap Day Friend",
			Birthdate:    time.Date(2000, time.February, 29, 0, 0, 0, 0, time.UTC),
			Timezone:     "Pacific/Auckland",
			ContactLinks: []string{},
			Wishlists:    []string{},
			Events:       []uuid.UUID{},
		},
		{
			ID:           uuid.MustParse("66666666-6666-6666-6666-666666666666"),
			Name:         "Bob's Contact",
			Birthdate:    time.Date(2000, time.March, 10, 0, 0, 0, 0, time.UTC),
			Timezone:     "Europe/Lisbon",
			ContactLinks: []string{},
			Wishlists:    []string{},
			Events:       []uuid.UUID{},
		},
	}
}

func getMockUsersWithBirthdays() []models.User {
	users := getMockUsersWithSessions()
	users[0].Birthdays = []string{
		"44444444-4444-4444-4444-444444444444",
		"55555555-5555-5555-5555-555555555555",
	}
	users[1].Birthdays = []string{
		"66666666-6666-6666-6666-666666666666",
	}
	return users
}

func getMockUsers() []models.User {
	return []models.User{
		{
			ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:      "Alice Nguyen",
			Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
			Email:     "alice.nguyen@example.com",
			Timezone:  "America/New_York",
			Birthdays: nil,
			Sessions:  nil,
		},
		{
			ID:        uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:      "Bob Mercado",
			Password:  "$2a$10$X7trdoO6c2Ym0XWR1vNXVeWbnqZ6yNlTeqrMc0BpeDBMdCvQRtszi",
			Email:     "bob.mercado@example.com",
			Timezone:  "Europe/Lisbon",
			Birthdays: nil,
			Sessions:  nil,
		},
		{
			ID:        uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			Name:      "Chiemi Tanaka",
			Password:  "$2a$10$cZ4gVKFmqkhRRRUnA5MTYuQcwBTmEQT3wjQMxjOkCEwp0GT2EfZ1O",
			Email:     "chiemi.tanaka@example.co.jp",
			Timezone:  "Asia/Tokyo",
			Birthdays: nil,
			Sessions:  nil,
		},
	}
}

func getMockUsersWithSessions() []models.User {
	return []models.User{
		{
			ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:      "Alice Nguyen",
			Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
			Email:     "alice.nguyen@example.com",
			Timezone:  "America/New_York",
			Birthdays: nil,
			Sessions: []models.Session{
				{
					ID:          uuid.MustParse("33333333-3333-3333-3333-333333333333"),
					AccessToken: utilities.ToHash("8e1jgIYXjcoHvNDFmvKff0RPQBFGtZZl4nqK4u_OkRI"),
					CreatedAt:   time.Date(2030, time.January, 15, 12, 0, 0, 0, time.UTC),
					ExpiresAt:   time.Date(2030, time.January, 15, 12, 0, 0, 0, time.UTC).Add(90 * 24 * time.Hour),
				},
			},
		},
		{
			ID:        uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:      "Bob Mercado",
			Password:  "$2a$10$X7trdoO6c2Ym0XWR1vNXVeWbnqZ6yNlTeqrMc0BpeDBMdCvQRtszi",
			Email:     "bob.mercado@example.com",
			Timezone:  "Europe/Lisbon",
			Birthdays: nil,
			Sessions: []models.Session{
				{
					ID:          uuid.MustParse("d3b07384-0c23-4e7a-8c9b-2c4a5f6e7b9c"),
					AccessToken: utilities.ToHash("Z_TIeQlFa-KqzusUn5div552Rt-cND7vlBOB8fFbR4s"),
					CreatedAt:   time.Date(2030, time.January, 15, 12, 0, 0, 0, time.UTC),
					ExpiresAt:   time.Date(2030, time.January, 15, 12, 0, 0, 0, time.UTC).Add(90 * 24 * time.Hour),
				},
			},
		},
	}
}
