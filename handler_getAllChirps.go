package main

import (
	"net/http"

	"github.com/ajswetz/Chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) getAllChirpsHandler(w http.ResponseWriter, r *http.Request) {

	var dbChirps []database.Chirp
	var err error
	var userID uuid.UUID

	// Check for 'sort' query parameter
	sortOrder := r.URL.Query().Get("sort")
	switch sortOrder {
	case "desc":
		// Do nothing
	case "asc":
		// Do nothing
	default:
		// If no sort param was provided, or if an invalid param was provided, use "asc" as the default
		sortOrder = "asc"
	}

	// Check for 'author_id' query parameter
	authorID := r.URL.Query().Get("author_id")
	// If author_id is provided, parse into UUID value
	if authorID != "" {
		userID, err = uuid.Parse(authorID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
		}
	}

	switch {
	case sortOrder == "asc" && authorID == "":
		// Get all chirps in ascending order
		dbChirps, err = cfg.db.GetAllChirpsAsc(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps from the database", err)
			return
		}

	case sortOrder == "asc" && authorID != "":
		// Get only chirps for specified author in ascending order
		dbChirps, err = cfg.db.GetChirpsForGivenAuthorAsc(r.Context(), userID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps from the database", err)
			return
		}

	case sortOrder == "desc" && authorID == "":
		// Get all chirps in descending order
		dbChirps, err = cfg.db.GetAllChirpsDesc(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps from the database", err)
			return
		}

	case sortOrder == "desc" && authorID != "":
		// Get only chirps for specified author in descending order
		dbChirps, err = cfg.db.GetChirpsForGivenAuthorDesc(r.Context(), userID)
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
