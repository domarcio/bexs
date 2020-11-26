package handler

import "encoding/json"

type responseError struct {
	Message string `json:"message"`
}

func structToString(target interface{}) string {
	res, _ := json.Marshal(target)
	return string(res)
}

func stringToStruct(str []byte, target interface{}) interface{} {
	json.Unmarshal(str, target)
	return target
}
