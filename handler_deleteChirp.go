package main

import (
	"net/http"

	"github.com/ajswetz/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) deleteChirpHandler(w http.ResponseWriter, r *http.Request) {

	// Get JWT access token from Authorization header
	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing authentication token", err)
		return
	}

	// Validate JWT token
	userIDFromToken, err := auth.ValidateJWT(accessToken, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid authentication token", err)
		return
	}

	// Parse chirp ID from request path
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

	// Check whether requesting user is the author of the chirp
	if dbChirp.UserID != userIDFromToken {
		respondWithError(w, http.StatusForbidden, "User is not authorized to delete this chirp", err)
		return
	}

	// Delete the chirp
	err = cfg.db.DeleteChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "There was a problem deleting chirp", err)
		return
	}

	// Response with JSON
	respondWithJSON(w, http.StatusNoContent, nil)

}
