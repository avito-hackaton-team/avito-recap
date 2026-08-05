// Package entity contains recap domain entities.
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Recap struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Year        int
	Archetype   Archetype
	Slides      json.RawMessage
	GeneratedAt time.Time
}

type RecapCreation struct {
	ID      uuid.UUID
	Created bool
}
