// pkg/utils/auth_util.go

package utils

import (
	"strings"
)

// ExtractBearerToken extracts the Bearer token from the Authorization header.
func ExtractBearerToken(authorizationHeader string) string {
	splitToken := strings.Split(authorizationHeader, "Bearer ")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}