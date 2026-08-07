// Package profile implements profile business scenarios.
package profile

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

type profileRepository interface {
	List(ctx context.Context) ([]entity.Profile, error)

	GetByID(ctx context.Context, id uuid.UUID) (entity.Profile, error)
}

type profileService struct {
	profileRepository profileRepository
}

func NewProfileService(profileRepository profileRepository) *profileService {
	return &profileService{
		profileRepository: profileRepository,
	}
}

func (s *profileService) List(ctx context.Context) ([]entity.Profile, error) {
	profiles, err := s.profileRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list profiles: %w", err)
	}

	return profiles, nil
}
