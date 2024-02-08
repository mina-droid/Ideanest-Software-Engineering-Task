// pkg/utils/token_util.go

package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateToken generates a random token.
func GenerateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}