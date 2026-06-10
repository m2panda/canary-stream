package domain

import "context"

const (
	StatusActive  string = "active"
	StatusDisable string = "disabled"
	StatusPending string = "pending"
	StatusEx      string = "ex"
)

type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type StatusRepository interface {
	SelectAll(ctx context.Context) ([]Status, error)
}

type StatusUseCase interface {
	GetAll(ctx context.Context) ([]Status, error)
}
