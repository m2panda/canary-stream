package genre

import (
	"canary-stream/backend/internal/domain"
	"encoding/json"
	"net/http"
)

type getAllHandler struct {
	usecase domain.GenreUseCase
}

func (handler *getAllHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	genres, err := handler.usecase.GetAll(ctx)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(genres)
}

func NewGetAllHandler(usecase domain.GenreUseCase) *getAllHandler {
	return &getAllHandler{usecase: usecase}
}
