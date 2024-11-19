package main

import (
	"log"
	"slices"
	"strings"
)

func replaceProfanity(inputStr string) string {

	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}

	inputStrSlice := strings.Split(inputStr, " ")

	for i, word := range inputStrSlice {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			log.Printf("Replacing '%s' with '****'\n", word)
			inputStrSlice[i] = "****"
		}
	}

	cleanedStr := strings.Join(inputStrSlice, " ")

	return cleanedStr

}
