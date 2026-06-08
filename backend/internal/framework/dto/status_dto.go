package dto

import (
	"canary-stream/backend/internal/domain"
)

type StatusResponse struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (response *StatusResponse) Mapper(state domain.Status) {
	response.Name = state.Name
	response.Slug = state.Slug
}
