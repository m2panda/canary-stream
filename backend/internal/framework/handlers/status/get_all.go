package status

import (
	"canary-stream/backend/internal/domain"
	"encoding/json"
	"net/http"
)

type getAllHandler struct {
	usecase domain.StatusUseCase
}

func (handler *getAllHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	status, err := handler.usecase.GetAll(ctx)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(status)
}

func NewGetAllHandler(usecase domain.StatusUseCase) *getAllHandler {
	return &getAllHandler{usecase: usecase}
}
