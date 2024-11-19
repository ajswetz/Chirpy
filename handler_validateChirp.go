package main

import (
	"encoding/json"
	"net/http"
)

type Chirp struct {
	Body string `json:"body"`
}

type returnVals struct {
	CleanedBody string `json:"cleaned_body"`
}

func validateChirpHandler(resWriter http.ResponseWriter, req *http.Request) {

	var chirpJson Chirp
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&chirpJson)
	if err != nil {
		respondWithError(resWriter, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if len(chirpJson.Body) > 140 {
		// Invalid length - need to return an error
		respondWithError(resWriter, http.StatusBadRequest, "Chirp is too long", nil)
		return

	}

	// Replace profanity if found
	cleanChirp := replaceProfanity(chirpJson.Body)

	// else - return valid, cleaned response
	respondWithJSON(resWriter, http.StatusOK, returnVals{
		CleanedBody: cleanChirp,
	})

}
