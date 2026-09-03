package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `bson:"_id"`
	Name          string    `bson:"name"`
	Password      string    `bson:"password"`
	EncryptionKey string    `bson:"encryption_key"`
	Email         string    `bson:"email"`
	Timezone      string    `bson:"timezone"`
	Birthdays     []string  `bson:"birthdays"`
	Sessions      []Session `bson:"sessions"`
}

type Session struct {
	ID           uuid.UUID `bson:"_id"`
	AccessToken  string    `bson:"access_token"`
	RefreshToken string    `bson:"refresh_token"`
	CreatedAt    time.Time `bson:"created_at"`
	ExpiresAt    time.Time `bson:"expires_at"`
}
