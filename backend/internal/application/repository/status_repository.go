package repository

import (
	"canary-stream/backend/core"
	"canary-stream/backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type statusRepository struct {
	db *pgxpool.Pool
}

func (repository *statusRepository) SelectAll(ctx context.Context) ([]domain.Status, error) {
	rows, err := repository.db.Query(ctx, core.Queries["STATUS_GET_ALL"])

	if err != nil {
		return nil, fmt.Errorf("failed to query status data: %w", err)
	}

	defer rows.Close()

	var status []domain.Status

	for rows.Next() {
		var row domain.Status

		err = rows.Scan(
			&row.Name,
			&row.Slug,
		)

		if err != nil {
			return nil, fmt.Errorf("Failed to scan genre row: %w", err)
		}

		status = append(status, row)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error during row iteration: %w", err)
	}

	return status, nil
}

func NewStatusRepository(db *pgxpool.Pool) domain.StatusRepository {
	return &statusRepository{db: db}
}
