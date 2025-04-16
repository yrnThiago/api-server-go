package utils

type ErrorInfo struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (e *ErrorInfo) Error() string {
	return e.Message
}

func NewErrorInfo(name string, msg string) *ErrorInfo {
	return &ErrorInfo{
		Name:    name,
		Message: msg,
	}
}

func GetInternalError() *ErrorInfo {
	return &ErrorInfo{
		Name:    "internal server error",
		Message: "internal server error",
	}
}
