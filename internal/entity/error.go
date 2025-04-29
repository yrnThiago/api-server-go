package entity

import (
	"strings"
)

const (
	ErrValidationName      = "VALIDATION_ERROR"
	ErrInvalidJwtTokenName = "JWT_INVALID_TOKEN"
	ErrRecordNotFoundName  = "RECORD_NOT_FOUND"
	ErrInternalServerName  = "INTERNAL_SERVER_ERROR"
	ErrRateLimitName       = "RATE_LIMIT_ERROR"

	ErrValidationMsg       = "invalid body request"
	ErrRecordNotFoundMsg   = "id not found"
	ErrInternalServerMsg   = "something went wrong"
	ErrRateLimitMsg        = "too many requests"
	ErrUnauthorizedMsg     = "unauthorized"
	ErrWrongCredentialsMsg = "wrong credentials"
)

type ErrorInfo struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Errors  []ValidationError
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorInfo) Error() string {
	return e.Message
}

func (e *ErrorInfo) GetLowerName() string {
	return strings.ToLower(strings.ReplaceAll(e.Name, "_", " "))
}

func NewErrorInfo(name, msg string, errors []ValidationError) *ErrorInfo {
	return &ErrorInfo{
		Name:    name,
		Message: msg,
		Errors:  errors,
	}
}

func GetValidationError(validationErrors []ValidationError) *ErrorInfo {
	return NewErrorInfo(ErrValidationName, ErrValidationMsg, validationErrors)
}

func GetInvalidJwtTokenError() *ErrorInfo {
	return NewErrorInfo(ErrInvalidJwtTokenName, ErrUnauthorizedMsg, nil)
}

func GetInternalServerError() *ErrorInfo {
	return NewErrorInfo(ErrInternalServerName, ErrInternalServerMsg, nil)
}

func GetNotFoundError() *ErrorInfo {
	return NewErrorInfo(ErrRecordNotFoundName, ErrRecordNotFoundMsg, nil)
}

func GetRateLimitError() *ErrorInfo {
	return NewErrorInfo(ErrRateLimitName, ErrRateLimitMsg, nil)
}

func AsErrorInfo(err error) *ErrorInfo {
	errorInfo, ok := err.(*ErrorInfo)
	if !ok {
		return GetInternalServerError()
	}

	return errorInfo

}
