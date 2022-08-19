package api

import (
	"fmt"
	"net/http"
)

// CreateRoutes Registers every route on the api, starting with /api
func CreateRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	// Main API route
	http.HandleFunc("/api/", printServerStarted)

	HelloRoute()
}

func printServerStarted(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Auth server is ON!")
}
