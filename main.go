package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

func main() {

	// Create new HTTP request multiplexer
	mux := http.NewServeMux()

	// Create state tracker
	srvState := &apiConfig{
		fileserverHits: atomic.Int32{},
	}

	/// APP ///

	// Register file server handler wrapped by middlewareMetricsInc() to track number of hits
	fileServerHandler := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", srvState.middlewareMetricsInc(fileServerHandler))

	/// API ///

	// Register server readiness handler
	mux.HandleFunc("GET /api/healthz", readinessHandler)

	/// ADMIN ///

	// Register metrics handler
	mux.HandleFunc("GET /admin/metrics", srvState.metricsHandler)

	// Register reset metrics handler
	mux.HandleFunc("POST /admin/reset", srvState.resetMetricsHandler)

	// Define http.Server listening on port 8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start listening on port 8080 and serving responses
	log.Println("Starting Chirpy . . . ")
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
	}

}
