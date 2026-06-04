package domain

import "context"

type Status struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type StatusRepository interface {
	SelectAll(ctx context.Context) ([]Status, error)
}

type StatusUseCase interface {
	GetAll(ctx context.Context) ([]Status, error)
}
