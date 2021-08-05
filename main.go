package main

import (
	"log"
	"net/http"
	"sing3demons/go-rest-api/routes"
	"sing3demons/go-rest-api/utils"

	"github.com/gorilla/mux"
)

const port = "8080"

type Map map[string]interface{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		utils.Json(w, http.StatusOK)(Map{"msg": "Hello"})
	})

	routes.Serve(router)

	log.Println("Server listening on port: ", port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
