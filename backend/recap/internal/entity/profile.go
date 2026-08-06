package entity

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID           uuid.UUID
	Name         string
	Surname      string
	AvatarURL    *string
	Hint         string
	RegisteredAt time.Time
}
