package api

import (
	"fmt"
	"net/http"
)

// CreateRoutes Registers every route on the api, starting with /api
func CreateRoutes() {
	// Root route
	http.HandleFunc("/", printRootRouteWorking)

	// Main API route
	http.HandleFunc("/api/", printMainRouteWorking)

	// Users API route
	UsersRoute()

	// Dummy route to serve as a playground for test implementations
	HelloRoute()
}

func printMainRouteWorking(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Auth server is ON!")
}

func printRootRouteWorking(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
