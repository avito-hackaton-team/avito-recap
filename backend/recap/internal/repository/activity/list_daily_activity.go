package activity

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

func (r *Repository) ListDailyActivity(
	ctx context.Context,
	userID uuid.UUID,
	period entity.Period,
) ([]entity.DayActivity, error) {
	if !period.Valid() {
		return nil, fmt.Errorf(
			"list daily activity: invalid period [%v, %v)",
			period.From,
			period.To,
		)
	}

	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	const query = `
		WITH activity_events AS (
			SELECT v.viewed_at AS occurred_at
			FROM recap.views AS v
			WHERE v.user_id = $1
			  AND v.viewed_at >= $2
			  AND v.viewed_at < $3

			UNION ALL

			SELECT f.created_at AS occurred_at
			FROM recap.favorites AS f
			WHERE f.user_id = $1
			  AND f.created_at >= $2
			  AND f.created_at < $3

			UNION ALL

			SELECT d.completed_at AS occurred_at
			FROM recap.deals AS d
			WHERE d.buyer_id = $1
			  AND d.completed_at IS NOT NULL
			  AND d.completed_at >= $2
			  AND d.completed_at < $3

			UNION ALL

			SELECT d.completed_at AS occurred_at
			FROM recap.deals AS d
			JOIN recap.listings AS l ON l.id = d.listing_id
			WHERE l.seller_id = $1
			  AND d.completed_at IS NOT NULL
			  AND d.completed_at >= $2
			  AND d.completed_at < $3

			UNION ALL

			SELECT m.created_at AS occurred_at
			FROM recap.messages AS m
			WHERE m.buyer_id = $1
			  AND m.created_at >= $2
			  AND m.created_at < $3

			UNION ALL

			SELECT m.created_at AS occurred_at
			FROM recap.messages AS m
			WHERE m.seller_id = $1
			  AND m.created_at >= $2
			  AND m.created_at < $3

			UNION ALL

			SELECT l.created_at AS occurred_at
			FROM recap.listings AS l
			WHERE l.seller_id = $1
			  AND l.created_at >= $2
			  AND l.created_at < $3
		)
		SELECT
			(occurred_at AT TIME ZONE 'UTC')::date AS day,
			COUNT(*) AS actions
		FROM activity_events
		GROUP BY day
		ORDER BY day
	`

	rows, err := r.pool.Query(ctx, query, userID, period.From, period.To)
	if err != nil {
		return nil, fmt.Errorf("list daily activity: %w", err)
	}
	defer rows.Close()

	result := make([]entity.DayActivity, 0)
	for rows.Next() {
		var model dailyActivityModel
		if err := model.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan daily activity: %w", err)
		}

		result = append(result, entity.DayActivity{
			Date:    model.Day,
			Actions: model.Actions,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate daily activities: %w", err)
	}

	return result, nil
}
