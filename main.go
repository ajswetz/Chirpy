package main

import (
	"fmt"
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

	// Register file server handler wrapped by middlewareMetricsInc() to track number of hits
	fileServerHandler := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", srvState.middlewareMetricsInc(fileServerHandler))

	// Register server readiness handler
	mux.HandleFunc("/healthz", readinessHandler)

	// Register metrics handler
	mux.HandleFunc("/metrics", srvState.metricsHandler)

	// Register reset metrics handler
	mux.HandleFunc("/reset", srvState.resetMetricsHandler)

	// Define http.Server listening on port 8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start listening on port 8080 and serving responses
	fmt.Println("Starting Chirpy . . . ")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}

}
