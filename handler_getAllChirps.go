package main

import (
	"net/http"
)

func (cfg *apiConfig) getAllChirpsHandler(w http.ResponseWriter, r *http.Request) {

	// Send query to database to get all chirps
	dbChirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps from the database", err)
		return
	}

	// Map database.Chirp struct to main.Chirp struct for JSON key control
	allChirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		allChirps = append(allChirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserId:    dbChirp.UserID,
		})
	}

	// Response with JSON
	respondWithJSON(w, http.StatusOK, allChirps)

}
