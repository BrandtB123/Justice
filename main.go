package main

import (
	"fmt"
	"justice/models"
	"justice/router"
	"net/http"
)

func main() {
	models.NewDB()
	go func() {
		r := router.Router()
		err := http.ListenAndServe(":9000", r)
		if err != nil {
			fmt.Println("Error establishing connection:", err)
		}
	}()
}
