package status

import (
	"canary-stream/backend/internal/domain"
	"canary-stream/backend/internal/framework/dto"
	"encoding/json"
	"net/http"
)

type getAllHandler struct {
	usecase domain.StatusUseCase
}

/**
 * Handler to get complete status
 * registers as dictionary information
 */
func (handler *getAllHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var data []dto.StatusResponse

	ctx := request.Context()
	status, err := handler.usecase.GetAll(ctx)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, state := range status {
		var object dto.StatusResponse

		(&object).Mapper(state)

		data = append(data, object)
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func NewGetAllHandler(usecase domain.StatusUseCase) *getAllHandler {
	return &getAllHandler{usecase: usecase}
}
