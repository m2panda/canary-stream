package user

import (
	"bytes"
	"canary-stream/backend/internal/domain"
	"canary-stream/backend/internal/framework/dto"
	"canary-stream/backend/internal/framework/i18n"
	"canary-stream/backend/internal/framework/validation"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type createRegisterHandler struct {
	usecase domain.UserUseCase
}

func (handler *createRegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var user *dto.UserRegisterBody

	apiResponse := &dto.ApiResponse[bool]{Success: false}
	responseError := &dto.ErrorDetail{Message: i18n.DefaultError}

	/**
	 * If process is successful or not
	 * response always will be in json format
	 */
	response.Header().Set("Content-Type", "application/json")

	langHeader := request.Header.Get("Accept-Language")

	if validation.Validator == nil {
		slog.Error("Error validator is not available",
			"event", "validator.status",
			"handler", "user.create_register",
		)

		if message, err := i18n.Translate(langHeader, i18n.ErrServerValidatorNotAvailable, nil); err == nil {
			responseError.Message = message
		}

		apiResponse.Error = responseError

		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(apiResponse)
		return
	}

	bodyBytes, err := io.ReadAll(request.Body)

	if err != nil {
		slog.Error("Error reading request body",
			"event", "validator.read_body",
			"handler", "user.create_register",
			"error", err,
		)

		response.WriteHeader(http.StatusInternalServerError)

		if message, errT := i18n.Translate(langHeader, i18n.ErrServerReadingRequestBody, nil); errT == nil {
			responseError.Message = message
		}

		apiResponse.Error = responseError
		json.NewEncoder(response).Encode(apiResponse)
		return
	}

	request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&user); err != nil {
		slog.Error("Error parsing request body to valid format",
			"event", "validator.parse_body",
			"handler", "user.create_register",
			"body", string(bodyBytes),
			"error", err,
		)

		if message, errT := i18n.Translate(langHeader, i18n.ErrRequestParseBody, nil); errT == nil {
			responseError.Message = message
		}

		apiResponse.Error = responseError

		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(apiResponse)
		return
	}

	if err := (*validation.Validator).Struct(user); err != nil {
		slog.Error("Error validating request params in fields",
			"event", "validator.param_validation",
			"handler", "user.create_register",
			"body", string(bodyBytes),
			"error", err,
		)

		details := make(map[string]interface{})

		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, field := range errors {
				fieldName := strings.ToLower(field.Field())

				if _, exist := details[fieldName]; !exist {
					details[fieldName] = []dto.ValidationDetails{}
				}

				detail := &dto.ValidationDetails{
					Issue:    field.Tag(),
					Expected: field.Param(),
					Value:    fmt.Sprintf("%v", field.Value()),
					Fallback: i18n.DefaultError,
				}

				templateData := map[string]interface{}{
					"Field": field.Field(),
					"Param": field.Param(),
				}

				switch field.Tag() {
				case "alphanum":
					if message, err := i18n.Translate(langHeader, i18n.ErrRequestValidatorAlphanum, templateData); err == nil {
						detail.Fallback = message
					}
				case "min":
					detail.Value = strconv.Itoa(len(detail.Value))

					if message, err := i18n.Translate(langHeader, i18n.ErrRequestValidatorMin, templateData); err == nil {
						detail.Fallback = message
					}
				case "max":
					detail.Value = strconv.Itoa(len(detail.Value))

					if message, err := i18n.Translate(langHeader, i18n.ErrRequestValidatorMax, templateData); err == nil {
						detail.Fallback = message
					}
				case "securepassword":
					if message, err := i18n.Translate(langHeader, i18n.ErrRequestValidatorSecurepassword, templateData); err == nil {
						detail.Fallback = message
					}
				default:
					if message, err := i18n.Translate(langHeader, i18n.ErrRequestValidatorDefault, templateData); err == nil {
						detail.Fallback = message
					}
				}

				if value, allowed := details[fieldName].([]dto.ValidationDetails); allowed {
					details[fieldName] = append(value, *detail)
				}
			}
		}

		if message, err := i18n.Translate(langHeader, i18n.ErrRequestValidatorGeneral, nil); err == nil {
			responseError.Message = message
		}

		responseError.Details = details

		apiResponse.Error = responseError

		response.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(response).Encode(apiResponse)
		return
	}

	ctx := request.Context()

	success, err := handler.usecase.CreateRegister(ctx, user.Username, user.Password)

	if err != nil || !success {
		slog.Error("Error register new user",
			"event", "user.create_register",
			"handler", "user.create_register",
			"success", success,
			"error", err,
		)

		if message, errT := i18n.Translate(langHeader, i18n.ErrUserCreateRegister, nil); errT == nil {
			responseError.Message = message
		}

		apiResponse.Error = responseError

		response.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(response).Encode(apiResponse)
		return
	}

	apiResponse.Success = true
	json.NewEncoder(response).Encode(apiResponse)
}

func NewCreateRegisterHandler(usecase domain.UserUseCase) *createRegisterHandler {
	return &createRegisterHandler{usecase: usecase}
}
