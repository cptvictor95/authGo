package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		myName := "world!"
		fmt.Fprintf(w, "Hello, %s", html.EscapeString(myName))
	})

	http.ListenAndServe("localhost:4200", nil)
}