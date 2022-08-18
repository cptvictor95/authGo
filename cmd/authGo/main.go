package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Start the server, listen and log
	addr := "localhost:4200"
	log.Printf("Server is up and running on http://%s/", addr)
	log.Fatal((http.ListenAndServe(addr, nil)))

	// Main API route
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Auth server is ON!")
	})

	// First hello world route
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	myName := "Victor!"
	// 	fmt.Fprintf(w, "Hello, %s", html.EscapeString(myName))
	// })

	// Second hello world on another thread
	// go func ()  {
	// 	resp, err := http.Get("http://localhost:4200/hello")

	// 	body, err := ioutil.ReadAll(resp.Body)

	// 	if err != nil {
	// 		fmt.Println(err)

	// 		return
	// 	}

	// 	fmt.Println(string(body))
		
	// 	body := make([]byte, 420)
		
	// 	bodySize, err := resp.Body.Read(body)
	// 	fmt.Println(resp.Status)
	// 	fmt.Println(string(body[:bodySize]))
	// 	fmt.Println(err)
	// }()

	
}