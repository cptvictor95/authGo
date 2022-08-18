package api

import (
	"fmt"
	"net/http"
)

// CreateRoutes Registers every route on the api, starting with /api
func CreateRoutes() {
	// Main API route
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Auth server is ON!")
	})
}