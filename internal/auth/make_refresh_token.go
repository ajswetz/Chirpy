package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// Generate a random 256-bit (32-byte) hex-encoded string
func MakeRefreshToken() (string, error) {

	// Use crypto/rand.Read() to generate 32 bytes (256 bits) of random data
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("unable to generate random bytes using crypto/rand.Read() function")
	}

	// hex.EncodeToString to convert the random data to a hex string
	hexString := hex.EncodeToString(bytes)

	return hexString, nil
}
