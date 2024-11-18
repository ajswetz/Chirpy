package main

import (
	"net/http"
	"strconv"
)

func (cfg *apiConfig) metricsHandler(resWriter http.ResponseWriter, _ *http.Request) {

	resWriter.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	resWriter.WriteHeader(200)

	// Get current file server hits from server state, convert to string, build full body text
	pageHits := cfg.fileserverHits.Load()
	hitsStr := strconv.Itoa(int(pageHits))
	resText := "Hits: " + hitsStr

	// Write response text using .Write() method
	resWriter.Write([]byte(resText))

}

func (cfg *apiConfig) resetMetricsHandler(resWriter http.ResponseWriter, _ *http.Request) {

	resWriter.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	resWriter.WriteHeader(200)

	// Reset current server hits metric to 0
	cfg.fileserverHits.Store(int32(0))

	// Write response text using .Write() method
	resWriter.Write([]byte("Page hit metric has been reset to 0."))

}
