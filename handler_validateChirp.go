package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Chirp struct {
	Body string `json:"body"`
}

type respValid struct {
	Valid bool `json:"valid"`
}

type respError struct {
	Error string `json:"error"`
}

func validateChirpHandler(resWriter http.ResponseWriter, req *http.Request) {

	resWriter.Header().Add("Content-Type:", "application/json")

	var chirpJson Chirp
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&chirpJson)
	if err != nil {
		log.Printf("An error occured decoding json request body\n")
		log.Printf("Error: %v\n", err)
	}

	if len(chirpJson.Body) > 140 {
		// Invalid length - need to return an error
		respErr := respError{
			Error: "Chirp is too long",
		}
		resWriter.WriteHeader(400)
		errJson, err := json.Marshal(respErr)
		if err != nil {
			log.Printf("An error occured marshaling struct to json\n")
			log.Printf("Error: %v\n", err)
		}
		resWriter.Write(errJson)
	} else {

		// else - return "valid" response
		respValid := respValid{
			Valid: true,
		}
		validJson, err := json.Marshal(respValid)
		if err != nil {
			log.Printf("An error occured marshaling struct to json\n")
			log.Printf("Error: %v\n", err)
		}
		resWriter.WriteHeader(200)
		resWriter.Write(validJson)

	}

}
