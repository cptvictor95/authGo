package api

import (
	"fmt"
	"net/http"
)

// HelloRoute Registers the dummy route for testing and playing with new knowledge
func HelloRoute() {
	http.HandleFunc("/api/hello/", func(w http.ResponseWriter, r *http.Request) {
		name := "world"
		fmt.Fprintf(w, "Hello, %s!", name)
	})
}