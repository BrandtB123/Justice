package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting the application...")

	// Perform configuration, database connections, etc.

	// router := setupRouter()

	// port := 8080
	// fmt.Printf("Listening on port %d...\n", port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// func setupRouter() *http.ServeMux {
// 	router := http.NewServeMux()
// 	// Define your routes here
// 	router.HandleFunc("/api/v1/hello", handleHello)
// 	return router
// }

// func handleHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }
