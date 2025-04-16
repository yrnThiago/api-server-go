package utils

type ErrorInfo struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func NewErrorInfo(statusCode int, err string) *ErrorInfo {
	return &ErrorInfo{
		StatusCode: statusCode,
		Error:      err,
	}
}
