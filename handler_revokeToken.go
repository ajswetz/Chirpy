package main

import (
	"net/http"

	"github.com/ajswetz/Chirpy/internal/auth"
)

func (cfg *apiConfig) revokeTokenHandler(w http.ResponseWriter, r *http.Request) {

	// Get refresh token from Authorization header
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing 'Authorizaton' header value", err)
		return
	}

	// Revoke token
	cfg.db.RevokeRefreshToken(r.Context(), refreshToken)

	// Refresh token successfully revoked - issue response
	respondWithJSON(w, http.StatusNoContent, nil)

}
