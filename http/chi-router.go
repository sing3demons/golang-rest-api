package router

import "net/http"

type chaiRouter struct {
}

func NewChoRouter() Router { return &chaiRouter{} }

func (*chaiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request))
func (*chaiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request))
func (*chaiRouter) SERVE(port string)
