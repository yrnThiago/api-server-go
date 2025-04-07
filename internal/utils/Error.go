package utils

type ErrorInfo struct {
	StatusCode int
	Message    string
}

func NewErrorInfo(statusCode int, message string) *ErrorInfo {
	return &ErrorInfo{
		StatusCode: statusCode,
		Message:    message,
	}
}
