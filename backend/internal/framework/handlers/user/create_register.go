package user

import (
	"canary-stream/backend/internal/domain"
	"canary-stream/backend/internal/framework/dto"
	"encoding/json"
	"net/http"
)

type createRegisterHandler struct {
	usecase domain.UserUseCase
}

func (handler *createRegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var user dto.UserRegisterBody
	var apiResponse dto.ApiResponse[bool]
	var responseError dto.ErrorDetail

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		responseError = dto.ErrorDetail{
			Code:    400,
			Message: "Data no valid",
		}

		apiResponse.Error = &responseError

		json.NewEncoder(response).Encode(apiResponse)
		return
	}

	ctx := request.Context()

	success, err := handler.usecase.CreateRegister(ctx, user.Username, user.Password)

	if err != nil && !success {
		responseError = dto.ErrorDetail{
			Code:    500,
			Message: "Server error on register",
		}

		apiResponse.Error = &responseError

		json.NewEncoder(response).Encode(apiResponse)
		return
	}

	apiResponse.Success = true
	json.NewEncoder(response).Encode(apiResponse)
}

func NewCreateRegisterHandler(usecase domain.UserUseCase) *createRegisterHandler {
	return &createRegisterHandler{usecase: usecase}
}
