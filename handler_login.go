package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ajswetz/Chirpy/internal/auth"
	"github.com/ajswetz/Chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type responseVals struct {
		ID           uuid.UUID `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Email        string    `json:"email"`
		IsChirpyRed  bool      `json:"is_chirpy_red"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}

	// Decode JSON from request
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	// Attempt to get user from database by email
	dbUser, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	// Check password hash
	err = auth.CheckPasswordHash(params.Password, dbUser.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	// Generate JSON Web Token
	expirationTime := time.Duration(time.Hour)
	newToken, err := auth.MakeJWT(dbUser.ID, cfg.secret, expirationTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to complete login at this time. Please try again later.", err)
		return
	}

	// Generate Refresh Token
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to complete login at this time. Please try again later.", err)
		return
	}

	// Store refresh token in the database
	createRefreshTokenParams := database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    dbUser.ID,
		ExpiresAt: time.Now().Add(time.Hour * 1440), // refresh token expires 60 days after creation
	}
	err = cfg.db.CreateRefreshToken(r.Context(), createRefreshTokenParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to complete login at this time. Please try again later.", err)
		return
	}

	// Passwords match - respond with JSON
	response := responseVals{
		ID:           dbUser.ID,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		Email:        dbUser.Email,
		IsChirpyRed:  dbUser.IsChirpyRed,
		Token:        newToken,
		RefreshToken: refreshToken,
	}
	respondWithJSON(w, http.StatusOK, response)

}
