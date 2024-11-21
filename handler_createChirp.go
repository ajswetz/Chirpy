package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ajswetz/Chirpy/internal/auth"
	"github.com/ajswetz/Chirpy/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserId    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) createChirpHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}

	// Get JWT Bearer Token
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing 'Authorizaton' header value", err)
		return
	}

	// Validate JWT token
	userIDFromToken, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid authentication token", err)
		return
	}

	// Decode JSON from request
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	// Validate chirp length
	const maxChirpLen = 140
	if len(params.Body) > maxChirpLen {
		// Invalid length - need to return an error
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	// Replace profanity if found
	params.Body = replaceProfanity(params.Body)

	// Send query to database to create new chirp
	newChirpParams := database.CreateChirpParams{
		Body:   params.Body,
		UserID: userIDFromToken,
	}
	dbChirp, err := cfg.db.CreateChirp(r.Context(), newChirpParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp in the database", err)
		return
	}

	// Map database.Chirp struct to main.Chirp struct for JSON key control
	newChirp := Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserId:    dbChirp.UserID,
	}

	// Response with JSON
	respondWithJSON(w, http.StatusCreated, newChirp)

}
