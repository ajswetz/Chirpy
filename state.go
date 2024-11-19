package main

import (
	"net/http"
	"sync/atomic"

	"github.com/ajswetz/Chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Incrememt file server hits counter
		cfg.fileserverHits.Add(1)

		//Call next handler in the chain
		next.ServeHTTP(w, r)

	})
}
