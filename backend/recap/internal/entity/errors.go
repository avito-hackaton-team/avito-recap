package entity

import "errors"

var (
	ErrProfileIDRequired = errors.New("profile id is required")

	ErrProfileNotFound = errors.New("profile not found")

	ErrRecapIDRequired = errors.New("recap id is required")

	ErrRecapNotFound = errors.New("recap not found")

	ErrNotEnoughActivity = errors.New("not enough activity to build a recap")

	ErrSharedRecapTokenInvalid = errors.New("shared recap token is invalid")

	ErrSharedRecapNotFound = errors.New("shared recap not found")
)
