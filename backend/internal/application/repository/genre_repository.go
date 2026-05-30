package repository

import (
	"canary-stream/backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type genreRepository struct {
	db *pgxpool.Pool
}

func (repository *genreRepository) SelectAll(ctx context.Context) ([]domain.Genre, error) {
	query := `SELECT _id, name, alt_name, slug FROM genres`

	rows, err := repository.db.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to query genre data: %w", err)
	}

	defer rows.Close()

	var genres []domain.Genre

	for rows.Next() {
		var genre domain.Genre

		err = rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.AltName,
			&genre.Slug,
		)

		if err != nil {
			return nil, fmt.Errorf("Failed to scan genre row: %w", err)
		}

		genres = append(genres, genre)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error during row iteration: %w", err)
	}

	return genres, nil
}

func NewGenreRepository(db *pgxpool.Pool) domain.GenreRepository {
	return &genreRepository{db: db}
}
