package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	// Hash the password using the bcrypt.GenerateFromPassword function. Bcrypt is a secure hash function that is intended for use with passwords.

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	} else {
		return string(hashedPass), nil
	}

}
