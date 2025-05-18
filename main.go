package main

import (
	"net/http"
	"log"
	"sync/atomic"
	"fmt"
)

// Create a struct to hold the hit counter
type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {

	const port  = "8080"

	// instantiate from apiConfig
	apiCfg := &apiConfig{}

	// Create a new http.ServeMux
	mux := http.NewServeMux()

	appHandler := http.FileServer(http.Dir("."))

	// wrap the http.Fileserver with the middleware
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", appHandler)))

	// --- Metrics Endpoint ---
	// Step 5: Register the handler for the /metrics path.
	mux.HandleFunc("GET /metrics", apiCfg.handlerMetrics)

	// --- Reset Endpoint ---
	// Step 6: Create and register a handler on the /reset path.
	mux.HandleFunc("POST /reset", apiCfg.handlerReset)

	// add readines endpoint
	mux.HandleFunc("GET /healthz", healthzHandler)

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

// middleware method on *apiConfig
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	// This returned function is the actual middleware handler
	return http.HandlerFunc ( func(w http.ResponseWriter, r *http.Request){
		// increment the counter for every request that passes through the handler
		cfg.fileserverHits.Add(1)

		//call the next handler in the chain
		next.ServeHTTP(w,r)
	})
}

// handler that writes the number of requests
// this method is on *apiConfig and access fileserverHits

func (cfg *apiConfig) handlerMetrics (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	hits := cfg.fileserverHits.Load()
	responseText := fmt.Sprintf("Hits: %d", hits)
	w.Write([]byte(responseText))
}

// handler to reset the counter
func (cfg *apiConfig) handlerReset (w http.ResponseWriter, r *http.Request){
	cfg.fileserverHits.Store(0)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}

// healthzHandler is the handler for the /healthz endpoint
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}