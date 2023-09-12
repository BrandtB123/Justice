package rest

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// apis
	r.HandleFunc("/", Base)
	r.HandleFunc("/hello", Hello)
	r.HandleFunc("/hello/{name}", Name)

	// kube liveliness
	r.HandleFunc("/health", Health)
	r.HandleFunc("/readiness", Readiness)

	return r
}
