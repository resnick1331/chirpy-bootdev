package main

import (
	"net/http"
	"log"
)

func main() {

	const port  = "8080"
	// Create a new http.ServeMux
	mux := http.NewServeMux()

	// add readines endpoint
	mux.HandleFunc("/healthz", healthzHandler)

	// update fileserver path
	//assetHandler := http.FileServer(http.Dir("./assets"))
	//mux.Handle("/app/", http.StripPrefix("/app", assetHandler))

	appHandler := http.FileServer(http.Dir("."))
	mux.Handle("/app/", http.StripPrefix("/app", appHandler))


	//crete a file server
	//fs := http.FileServer(http.Dir("."))

	//handle request to root path
	//mux.Handle("/", fs)

	//create a new http.Server struct
	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", appHandler,port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

// healthzHandler is the handler for the /healthz endpoint
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}