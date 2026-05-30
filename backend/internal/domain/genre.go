package domain

import "context"

type Genre struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	AltName *string `json:"alt_name"`
	Slug    string  `json:"slug"`
}

type GenreRepository interface {
	SelectAll(ctx context.Context) ([]Genre, error)
}

type GenreUseCase interface {
	GetAll(ctx context.Context) ([]Genre, error)
}
