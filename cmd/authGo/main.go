package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func main() {
	// First hello world route
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		myName := "world!"
		fmt.Fprintf(w, "Hello, %s", html.EscapeString(myName))
	})

	// Second hello world on another thread
	go func ()  {
		resp, err := http.Get("http://localhost:4200/hello")

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(string(body))

		// body := make([]byte, 420)
		
		// bodySize, err := resp.Body.Read(body)
		// fmt.Println(resp.Status)
		// fmt.Println(string(body[:bodySize])) 
		// fmt.Println(err)
	}()

	http.ListenAndServe("localhost:4200", nil)
}