package rest

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", Base)
	r.HandleFunc("/hello", Hello)
	r.HandleFunc("/hello/{name}", Name)

	return r
}
