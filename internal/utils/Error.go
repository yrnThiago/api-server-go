package utils

type ErrorInfo struct {
	Name       string `json:"name"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func NewErrorInfo(name string, statusCode int, err string) *ErrorInfo {
	return &ErrorInfo{
		Name:       name,
		StatusCode: statusCode,
		Error:      err,
	}
}
