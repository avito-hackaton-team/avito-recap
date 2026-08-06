package entity

import "errors"

var (
	ErrProfileNotFound = errors.New("profile not found")

	ErrRecapNotFound = errors.New("recap not found")

	ErrNotEnoughActivity = errors.New("not enough activity to build a recap")
)
