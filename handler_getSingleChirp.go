package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getSingleChirpHandler(w http.ResponseWriter, r *http.Request) {

	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to decode requested chirp ID into valid UUID", err)
		return
	}

	// Send query to database to get chirp by ID
	dbChirp, err := cfg.db.GetSingleChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	// Map database.Chirp struct to main.Chirp struct for JSON key control
	singleChirp := Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserId:    dbChirp.UserID,
	}

	// Response with JSON
	respondWithJSON(w, http.StatusOK, singleChirp)

}
