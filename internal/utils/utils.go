package utils

import "encoding/json"

func ConvertJsonToStruct[T any](body []byte, s T) any {
	err := json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}

	return s
}
