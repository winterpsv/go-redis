package apperror

import "fmt"

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	AppError   error  `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Error)
}

func NewAppError(statusCode int, message string, err error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
		AppError:   err,
	}
}
