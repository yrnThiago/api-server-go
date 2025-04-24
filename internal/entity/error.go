package entity

import "strings"

var (
	ErrProductNotFound = NewErrorInfo("RECORD_NOT_FOUND", "product id not found")
	ErrOrderNotFound   = NewErrorInfo("RECORD_NOT_FOUND", "order id not found")
	ErrUserNotFound    = NewErrorInfo("RECORD_NOT_FOUND", "user id not found")
	ErrRateLimit       = NewErrorInfo("RATE_LIMIT_ERROR", "too many requests")
	ErrInternalServer  = NewErrorInfo("INTERNAL_SERVER_ERROR", "something went wrong")

	ErrUnauthorizedMsg   = "unauthorized"
	ErrInternalServerMsg = "something went wrong"
)

type ErrorInfo struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (e *ErrorInfo) Error() string {
	return e.Message
}

func (e *ErrorInfo) GetLowerName() string {
	return strings.ToLower(strings.ReplaceAll(e.Name, "_", " "))
}

func NewErrorInfo(name, msg string) *ErrorInfo {
	return &ErrorInfo{
		Name:    name,
		Message: msg,
	}
}

func GetValidationError(msg string) *ErrorInfo {
	return &ErrorInfo{
		Name:    "VALIDATION_ERROR",
		Message: msg,
	}
}

func GetInvalidJwtTokenError(msg string) *ErrorInfo {
	return &ErrorInfo{
		Name:    "JWT_INVALID_TOKEN",
		Message: msg,
	}
}

func AsErrorInfo(err error) *ErrorInfo {
	errorInfo, ok := err.(*ErrorInfo)
	if !ok {
		return ErrInternalServer
	}

	return errorInfo

}
