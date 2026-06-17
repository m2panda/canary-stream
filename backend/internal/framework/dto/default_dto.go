package dto

type ValidationDetails struct {
	Issue    string `json:"issue"`
	Expected string `json:"expected"`
	Value    string `json:"value"`
	Fallback string `json:"fallback"`
}

type ErrorDetail struct {
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

type ApiResponse[T any] struct {
	Success bool         `json:"success"`
	Error   *ErrorDetail `json:"error"`
	Data    *T           `json:"data"`
}
