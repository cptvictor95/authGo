package main

import (
	"log"
	"net/http"

	"github.com/cptvictor95/authGo/internal/api"
)

func main() {
	// Register the main route /api with all subtree routes
	api.CreateRoutes()
	
	// Start the server with listener and logger
	bootstrap()
}

// Start the server, listen and log
func bootstrap() {
	addr := "localhost:4200"
	log.Printf("Server is up and running on http://%s/", addr)
	log.Fatal((http.ListenAndServe(addr, nil)))
}