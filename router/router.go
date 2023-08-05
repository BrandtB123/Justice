package router

import "net/http"

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/", nil)
	return mux
}
