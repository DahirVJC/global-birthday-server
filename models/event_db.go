package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID  `bson:"_id"`
	EventDateStart time.Time  `bson:"event_date_start"`
	EventDateEnd   *time.Time `bson:"event_date_end,omitempty"`
	RemindMeIn     *time.Time `bson:"remind_me_in,omitempty"`
	Description    *string    `bson:"description,omitempty"`
}
