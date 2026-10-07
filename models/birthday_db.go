package models

import (
	"time"

	"github.com/google/uuid"
)

type Birthday struct {
	ID           uuid.UUID   `bson:"_id"`
	Name         string      `bson:"name"`
	Birthdate    time.Time   `bson:"birthdate"`
	Timezone     string      `bson:"timezone"`
	RemindMeIn   *time.Time  `bson:"remind_me_in,omitempty"`
	ContactLinks []string    `bson:"contact_links,omitempty"`
	Wishlists    []string    `bson:"wishlists,omitempty"`
	Events       []uuid.UUID `bson:"events,omitempty"`
}
