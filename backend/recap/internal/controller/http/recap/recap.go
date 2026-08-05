// Package recap exposes recap delivery handlers.
package recap

import (
	"context"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type (
	recapService interface {
		Create(ctx context.Context, id uuid.UUID, year int) (uuid.UUID, error)
		Get(ctx context.Context, recapID uuid.UUID) (entity.Recap, error)
		// Share(ctx context.Context, id uuid.UUID) (entity.PublicRecap, error)
	}

	profileService interface {
		List(ctx context.Context) ([]entity.Profile, error)
	}
)

type recapServer struct {
	logger         *zap.Logger
	recapService   recapService
	profileService profileService
}

func NewRecapServer(logger *zap.Logger, recapService recapService, profileService profileService) *recapServer {
	return &recapServer{
		logger:         logger,
		recapService:   recapService,
		profileService: profileService,
	}
}
