package usecase

import (
	"canary-stream/backend/internal/domain"
	"context"
)

type genreUseCase struct {
	repository domain.GenreRepository
}

func (usecase *genreUseCase) GetAll(ctx context.Context) ([]domain.Genre, error) {
	return usecase.repository.SelectAll(ctx)
}

func NewGenreUseCase(repository domain.GenreRepository) domain.GenreUseCase {
	return &genreUseCase{repository: repository}
}
