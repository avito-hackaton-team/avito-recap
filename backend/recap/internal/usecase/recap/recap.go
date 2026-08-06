// Package recap implements recap business scenarios.
package recap

import (
	"context"

	"github.com/google/uuid"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

type ActivityRepository interface {
	GetActivityTotals(
		ctx context.Context,
		userID uuid.UUID,
		period entity.Period,
	) (entity.UserActivity, error)

	ListActivityByCategories(
		ctx context.Context,
		userID uuid.UUID,
		period entity.Period,
	) ([]entity.CategoryActivity, error)
}

type ProfileRepository interface {
	List(ctx context.Context) ([]entity.Profile, error)

	GetByID(ctx context.Context, id uuid.UUID) (entity.Profile, error)
}

type RecapRepository interface {
	Create(ctx context.Context, recap entity.Recap) (entity.RecapCreation, error)

	GetByID(ctx context.Context, id uuid.UUID) (entity.Recap, error)

	GetByProfileAndYear(
		ctx context.Context,
		profileID uuid.UUID,
		year int32,
	) (entity.Recap, error)
}
