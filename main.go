package main

import (
	"net/http"
	"log"
)

func main() {
	// Create a new http.ServeMux
	mux := http.NewServeMux()


	//create a new http.Server struct
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	log.Println("Server listening on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}