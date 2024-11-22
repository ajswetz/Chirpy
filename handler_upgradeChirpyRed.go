package main

import (
	"encoding/json"
	"net/http"

	"github.com/ajswetz/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) upgradeToChirpyRedHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	// Get API Key from header
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing 'Authorizaton' header value", err)
		return
	}

	// Validate API Key to ensure request is coming from Polka
	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Invalid API Key", err)
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

	// Don't care about any event other than 'user.upgraded' - respond 204
	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}

	// If we got this far, Event must be 'user.upgraded' - update user in database marking them as a Chirpy Red member
	err = cfg.db.SetChirpyRedTrue(r.Context(), params.Data.UserID)
	if err != nil {
		// Probably means user couldn't be found in the DB - respond 404
		respondWithError(w, http.StatusNotFound, "User not found", err)
	}

	// Else - respond 204 and empty response body
	respondWithJSON(w, http.StatusNoContent, nil)

}
