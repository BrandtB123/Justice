package main

import (
	"justice/internal/logger"
	"justice/internal/rest"
	"net/http"
)

func main() {
	r := rest.SetupRouter()
	logger.Init()
	logger.Logger.Info("test")
	http.ListenAndServe(":9000", r)
}
