// pkg/api/handlers/refresh_token_handler.go

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RefreshTokenHandler handles the refresh-token logic.
func RefreshTokenHandler(c *gin.Context) {
	// Parse JSON request body
	var refreshTokenData RefreshTokenData
	if err := c.ShouldBindJSON(&refreshTokenData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the refresh token (replace this with your actual refresh token validation logic)
	if isValidRefreshToken(refreshTokenData.RefreshToken) {
		// Generate a new access token and refresh token (replace this with your token generation logic)
		newAccessToken := generateAccessToken()
		newRefreshToken := generateRefreshToken()

		// Respond with success and new tokens
		c.JSON(http.StatusOK, gin.H{
			"message":        "Token refreshed successfully",
			"access_token":   newAccessToken,
			"refresh_token":  newRefreshToken,
		})
	} else {
		// Respond with refresh token validation failure
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
	}
}

// RefreshTokenData represents the request body for the refresh-token endpoint.
type RefreshTokenData struct {
	RefreshToken string `json:"refresh_token"`
}

// isValidRefreshToken checks if the provided refresh token is valid.
// Replace this function with your actual refresh token validation logic.
func isValidRefreshToken(refreshToken string) bool {
	// Your refresh token validation logic goes here
	// For demonstration purposes, always return true
	return true
}