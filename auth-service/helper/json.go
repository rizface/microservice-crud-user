package helper

import (
	"auth-service/model/response"
	"encoding/json"
	"net/http"
)

func JsonWriter(w http.ResponseWriter, code int, msg string, data map[string]interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response.Standard{
		Code:   code,
		Status: msg,
		Data:   data,
	})
}
