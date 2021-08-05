package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, statusCode int) func(v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)

	return func(v interface{}) error {
		return json.NewEncoder(w).Encode(v)
	}
}

func ToJson(w http.ResponseWriter, statusCode int) func(v interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	return func(v interface{}) error {
		result, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error marshalling the posts array"}`))
			return err
		}

		w.WriteHeader(statusCode)
		w.Write(result)
		return nil
	}
}

func ErrorToJson(w http.ResponseWriter, statusCode int) func(v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	return func(v interface{}) {
		w.WriteHeader(statusCode)
		w.Write([]byte(fmt.Sprintf(`{"error": "%v"}`, v)))
	}
}
