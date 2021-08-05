package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct {
}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router { return &chiRouter{} }

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi Http server running on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, chiDispatcher))
}
