package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {

	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no values associated with the 'Authorization' key in the headers received")
	}

	token, _ := strings.CutPrefix(authHeader, "Bearer")
	token = strings.TrimLeft(token, " ")

	return token, nil
}
