package models

type ApiError struct {
	Error string
}

func NewApiError(err error) ApiError {
	return ApiError{Error: err.Error()}
}
