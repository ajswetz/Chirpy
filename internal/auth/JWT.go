package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

	currentTimeUTC := time.Now().UTC()
	expiryTimeUTC := currentTimeUTC.Add(expiresIn)

	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(currentTimeUTC),
		ExpiresAt: jwt.NewNumericDate(expiryTimeUTC),
		Subject:   userID.String(),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedStr, err := newToken.SignedString([]byte(tokenSecret))
	if err != nil {
		log.Printf("An error occured signing JWT: %v\n", err)
		return "", err
	}

	// Else:
	return signedStr, nil

}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {

	claims := &jwt.RegisteredClaims{}

	// Parse the provided JWT ('tokenString') into a *jwt.Token. The function also fills in the `claims` struct with the data it extracts from the JWT
	// Uses an unnamed function as 3rd arg to return the 'tokenSecret', a.k.a. our secret key
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil })
	if err != nil {
		return uuid.UUID{}, err
	}

	// Get token subject from token Claims
	tokenSubject, err := token.Claims.GetSubject()
	if err != nil {
		log.Printf("An error occured getting Subject from token: %v\n", err)
		return uuid.UUID{}, err
	}

	// Convert stringified user id into uuid.UUID value
	userID, err := uuid.Parse(tokenSubject)
	if err != nil {
		log.Printf("An error occured converting token subject STRING to uuid.UUID: %v\n", err)
		return uuid.UUID{}, err
	}

	// Return validated user ID as uuid.UUID
	return userID, nil

}
