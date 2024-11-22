package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {

	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no values associated with the 'Authorization' key in the headers received")
	}

	apiKey, _ := strings.CutPrefix(authHeader, "ApiKey")
	apiKey = strings.TrimLeft(apiKey, " ")

	return apiKey, nil
}
