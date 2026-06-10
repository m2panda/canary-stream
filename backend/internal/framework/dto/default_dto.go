package dto

type ErrorDetail struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

type ApiResponse[T any] struct {
	Success bool         `json:"success"`
	Error   *ErrorDetail `json:"error"`
	Data    *T           `json:"data"`
}
