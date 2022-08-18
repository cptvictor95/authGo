package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cptvictor95/authGo/internal/api"
)

func main() {
	api.CreateRoutes()
	// Second hello world on another thread
	go func ()  {
		http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
			name := "world"
			fmt.Fprintf(w, "Hello, %s!", name)
		})
		resp, err := http.Get("http://localhost:4200/hello")

		// body, err := ioutil.ReadAll(resp.Body)

		// fmt.Println(string(body))
		
		body := make([]byte, 420)
		
		bodySize, err := resp.Body.Read(body)
		if err != nil {
			fmt.Println(err)

			return
		}
		
		fmt.Println(resp.Status)
		
		fmt.Println(string(body[:bodySize]))
		
		fmt.Println(err)
	}()

	bootstrap()
}

// Start the server, listen and log
func bootstrap() {
	addr := "localhost:4200"
	log.Printf("Server is up and running on http://%s/", addr)
	log.Fatal((http.ListenAndServe(addr, nil)))
}