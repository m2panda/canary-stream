package usecase

import (
	"canary-stream/backend/internal/domain"
	"context"
)

type statusUseCase struct {
	repository domain.StatusRepository
}

func (usecase *statusUseCase) GetAll(ctx context.Context) ([]domain.Status, error) {
	return usecase.repository.SelectAll(ctx)
}

func NewStatusUseCase(repository domain.StatusRepository) domain.StatusUseCase {
	return &statusUseCase{repository: repository}
}
