package api

import (
	"fmt"
	"net/http"
)

// HelloRoute Registers the dummy route as a playground for test implementations
func HelloRoute() {
	http.HandleFunc("/api/hello/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		
		name := query.Get("name")
		exists := query.Has("name")
		if exists && name == "" {
			fmt.Printf("Name parameter cannot be empty. Try again.")

			return
		}

		if !exists {
			name = "world"
		}
		
		fmt.Fprintf(w, "Hello, %s!", name)
	})
}