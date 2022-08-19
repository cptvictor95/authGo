package api

import (
	"fmt"
	"net/http"
)

// HelloRoute Registers the dummy route for testing and playing with new knowledge
func HelloRoute() {
	http.HandleFunc("/api/hello/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		
		name, exists := query["name"]

		if !exists || len(name) == 0 {
			fmt.Printf("Name parameter not found. Try again.")

			return
		}

		fmt.Fprintf(w, "Hello, %s!", name[0])
	})
}