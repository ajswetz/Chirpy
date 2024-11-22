package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ajswetz/Chirpy/internal/auth"
	"github.com/ajswetz/Chirpy/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Email       string    `json:"email"`
	IsChirpyRed bool      `json:"is_chirpy_red"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode JSON from request
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
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

	// Send query to database to create new user
	newUserParams := database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashedPass,
	}
	dbUser, err := cfg.db.CreateUser(r.Context(), newUserParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user in the database", err)
	}

	// Map database.user struct to main.User struct for JSON key control
	newUser := User{
		ID:          dbUser.ID,
		CreatedAt:   dbUser.CreatedAt,
		UpdatedAt:   dbUser.UpdatedAt,
		Email:       dbUser.Email,
		IsChirpyRed: dbUser.IsChirpyRed,
	}

	// Response with JSON
	respondWithJSON(w, http.StatusCreated, newUser)

}
