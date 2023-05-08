package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(password string) string {
	digest := sha256.Sum256([]byte(password))
	return hex.EncodeToString(digest[:])
}
