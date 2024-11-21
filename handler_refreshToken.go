package main

import (
	"net/http"
	"time"

	"github.com/ajswetz/Chirpy/internal/auth"
)

func (cfg *apiConfig) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	type responseVals struct {
		NewToken string `json:"token"`
	}

	// Get refresh token from Authorization header
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing 'Authorizaton' header value", err)
		return
	}

	// Look up refresh token in the database
	dbToken, err := cfg.db.GetRefreshToken(r.Context(), refreshToken)
	if err != nil {
		// refresh token not found in the database - respond with error
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token", err)
		return
	}
	if dbToken.RevokedAt.Valid {
		// refresh token has been revoked - respond with error
		respondWithError(w, http.StatusUnauthorized, "Expired refresh token", err)
		return
	}

	// Refresh token is valid and not expired - generate new JSON Web Token
	expirationTime := time.Duration(time.Hour)
	newToken, err := auth.MakeJWT(dbToken.UserID, cfg.secret, expirationTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to complete login at this time. Please try again later.", err)
		return
	}

	// Passwords match - respond with JSON
	response := responseVals{
		NewToken: newToken,
	}
	respondWithJSON(w, http.StatusOK, response)

}
