package user

import (
	"canary-stream/backend/core"
	"canary-stream/backend/internal/domain"
	"canary-stream/backend/internal/framework/dto"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type createRegisterHandler struct {
	usecase domain.UserUseCase
}

func (handler *createRegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if core.Validator == nil {
		return
	}

	defer request.Body.Close()

	var (
		user          dto.UserRegisterBody
		apiResponse   dto.ApiResponse[bool]
		responseError dto.ErrorDetail
	)

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

	if err := (*core.Validator).Struct(user); err != nil {
		var errors []string = []string{}
		var fieldErrMessage string

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, field := range validationErrors {
				switch field.Tag() {
				case "alphanum":
					fieldErrMessage = fmt.Sprintf("[%s]: El campo solo permite letras y números", field.Field())
				case "min":
					fieldErrMessage = fmt.Sprintf("[%s]: El campo requiere al menos %s caracteres", field.Field(), field.Param())
				case "max":
					fieldErrMessage = fmt.Sprintf("[%s]: El campo no puede tener más de %s caracteres", field.Field(), field.Param())
				case "securepassword":
					fieldErrMessage = fmt.Sprintf("[%s]: El campo no es válido, tiene que tener al menos una letra minúscula, una mayúscula, un número y un caracter especial", field.Field())
				default:
					fieldErrMessage = fmt.Sprintf("[%s]: El campo no es válido", field.Field())
				}

				errors = append(errors, fieldErrMessage)
			}
		}

		responseError = dto.ErrorDetail{
			Code:    400,
			Message: "Data validation error",
			Errors:  errors,
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
