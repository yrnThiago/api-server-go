package utils

type ErrorInfo struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewErrorInfo(statusCode int, message string) *ErrorInfo {
	return &ErrorInfo{
		StatusCode: statusCode,
		Message:    message,
	}
}
