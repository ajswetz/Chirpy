package main

import (
	"net/http"

	"github.com/ajswetz/Chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) getAllChirpsHandler(w http.ResponseWriter, r *http.Request) {

	var dbChirps []database.Chirp
	var err error

	// Check for 'author_id' query parameter
	authorID := r.URL.Query().Get("author_id")

	if authorID == "" {
		// No author_id query param provided - just get all chirps from DB
		dbChirps, err = cfg.db.GetAllChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps from the database", err)
			return
		}

	} else {
		// Only get Chirps with user_id that match provided author_id query
		userID, err := uuid.Parse(authorID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
		}
		dbChirps, err = cfg.db.GetChirpsForGivenAuthor(r.Context(), userID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps from the database", err)
			return
		}

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
