package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `bson:"_id"`
	Name      string    `bson:"name"`
	Password  string    `bson:"password"`
	Email     string    `bson:"email"`
	Timezone  string    `bson:"timezone"`
	Birthdays []string  `bson:"birthdays"`
	Sessions  []Session `bson:"sessions"`
}

type Session struct {
	ID          uuid.UUID `bson:"_id"`
	AccessToken string    `bson:"access_token"`
	CreatedAt   time.Time `bson:"created_at"`
	ExpiresAt   time.Time `bson:"expires_at"`
}

type Auth struct {
	ID        uuid.UUID `bson:"_id"`
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	Type      int16     `bson:"type"`
	CreatedAt time.Time `bson:"created_at"`
	ExpiresAt time.Time `bson:"expires_at"`
}
