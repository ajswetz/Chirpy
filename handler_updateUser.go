package main

import (
	"encoding/json"
	"net/http"

	"github.com/ajswetz/Chirpy/internal/auth"
	"github.com/ajswetz/Chirpy/internal/database"
)

func (cfg *apiConfig) updateUserHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

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

	// Decode JSON from request
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	// Hash password
	hashedPass, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "There was a problem hashing your password", err)
		return
	}

	// Send query to database to update user
	updateUserParams := database.UpdateUserParams{
		Email:          params.Email,
		HashedPassword: hashedPass,
		ID:             userIDFromToken,
	}
	dbUpdatedUser, err := cfg.db.UpdateUser(r.Context(), updateUserParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to update user account", err)
		return
	}

	// User account updated successfully - send JSON response
	updatedUser := User{
		ID:        dbUpdatedUser.ID,
		CreatedAt: dbUpdatedUser.CreatedAt,
		UpdatedAt: dbUpdatedUser.UpdatedAt,
		Email:     dbUpdatedUser.Email,
	}
	respondWithJSON(w, http.StatusOK, updatedUser)

}
