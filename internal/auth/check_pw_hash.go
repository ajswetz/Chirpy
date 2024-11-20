package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		// Password and hash do not match
		log.Println("Password and hash do not match")
		return err
	} else {
		// Password and hash DO match
		return nil
	}
}
