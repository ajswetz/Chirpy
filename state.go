package main

import (
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Incrememt file server hits counter
		cfg.fileserverHits.Add(1)

		//Call next handler in the chain
		next.ServeHTTP(w, r)

	})
}
