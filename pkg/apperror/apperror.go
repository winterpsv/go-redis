package apperror

import (
	"fmt"
	"log"
	"net/http"
)

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	AppError   error  `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Error)
}

func NewAppError(statusCode int, message string, err error) *AppError {
	log.Printf("[ERROR] %v: %v", statusCode, message)

	return &AppError{
		Message:    message,
		StatusCode: statusCode,
		AppError:   err,
	}
}

var (
	//Errors
	ErrExistsUser         = NewAppError(http.StatusConflict, "User already exists", nil)
	ErrCreateUser         = NewAppError(http.StatusInternalServerError, "Failed to create user", nil)
	ErrFindUser           = NewAppError(http.StatusInternalServerError, "Failed to find user", nil)
	ErrEditDeletedUser    = NewAppError(http.StatusNoContent, "User is deleted", nil)
	ErrUpdateUser         = NewAppError(http.StatusInternalServerError, "Failed to update user", nil)
	ErrAccess             = NewAppError(http.StatusUnauthorized, "Unauthorized user", nil)
	ErrPermission         = NewAppError(http.StatusForbidden, "Permission Denied", nil)
	ErrCreateToken        = NewAppError(http.StatusUnauthorized, "Failed to create token", nil)
	ErrAuthorizationToken = NewAppError(http.StatusUnauthorized, "Failed authorization token", nil)
	ErrParseToken         = NewAppError(http.StatusInternalServerError, "Failed to parse token", nil)
	ErrParseHeader        = NewAppError(http.StatusInternalServerError, "Failed to parse header", nil)
	ErrDecodeData         = NewAppError(http.StatusBadRequest, "Failed to decode data", nil)
	ErrValidation         = NewAppError(http.StatusUnprocessableEntity, "Validation failed", nil)

	//Success
	SuccessCreatedUser  = NewAppError(http.StatusCreated, "User created successfully", nil)
	SuccessUpdatedUser  = NewAppError(http.StatusOK, "User updated successfully", nil)
	SuccessCreatedToken = NewAppError(http.StatusCreated, "JWT token created successfully", nil)
)
