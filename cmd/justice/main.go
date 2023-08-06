package main

import (
	"justice/internal/rest"
	"net/http"
)

func main() {

	r := rest.SetupRouter()
	http.ListenAndServe(":9000", r)
}
