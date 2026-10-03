package utilities

import (
	"crypto/sha256"
	"encoding/hex"
)

func ToHash(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
