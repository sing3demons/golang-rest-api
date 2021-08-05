package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = "8080"

type Map map[string]interface{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		Json(w, http.StatusOK)(Map{"msg": "Hello"})
	})

	router.HandleFunc("/posts", getPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts", addPost).Methods(http.MethodPost)

	log.Println("Server listening on port: ", port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}

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
