package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(password string) string {
	digest := sha256.Sum256([]byte(password))
	return hex.EncodeToString(digest[:])
}

func ComparePasswords(inputPassword string, storedPassword string) bool {
	// Hash the input password
	inputHash := Sha256(inputPassword)

	// Compare the hashed input password with the stored password
	return inputHash == storedPassword
}