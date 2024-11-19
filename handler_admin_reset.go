package main

import "net/http"

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {

	// Check for "dev" platform environment variable
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Accessing this endpoint is forbidden", nil)
		return
	}

	// Set headers
	w.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	w.WriteHeader(200)

	/// RESET USERS ///
	err := cfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete users from database", err)
	}

	/// RESET METRICS ///
	cfg.fileserverHits.Store(int32(0))

	// Write response text using .Write() method
	w.Write([]byte("Page hit metric has been reset to 0. All users have been deleted from the database."))

}
