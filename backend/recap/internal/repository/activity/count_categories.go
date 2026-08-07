package activity

import (
	"context"
	"fmt"
)

func (r *Repository) CountCategories(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	const query = `
		SELECT COUNT(*)
		FROM recap.categories
	`

	var total int64
	if err := r.pool.QueryRow(ctx, query).Scan(&total); err != nil {
		return 0, fmt.Errorf("count categories: %w", err)
	}

	return total, nil
}
