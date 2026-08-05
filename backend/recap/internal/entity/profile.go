package entity

import "github.com/google/uuid"

type Profile struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Hint    string
}
