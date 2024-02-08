// pkg/api/handlers/create_organization_handler.go

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mina-droid/Ideanest-Software-Engineering-Task/pkg/utils"
)

// CreateOrganizationHandler handles the create-organization logic.
func CreateOrganizationHandler(c *gin.Context) {
	// Parse JSON request body
	var organizationData CreateOrganizationData
	if err := c.ShouldBindJSON(&organizationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate authentication token
	accessToken := utils.ExtractBearerToken(c.GetHeader("Authorization"))
	if !isValidAccessToken(accessToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid access token"})
		return
	}
	// Create organization
	organizationID := createOrganization(organizationData.Name, organizationData.Description)
	// Respond with success and organization ID
	c.JSON(http.StatusOK, gin.H{"organization_id": organizationID})
}

// CreateOrganizationData represents the request body for the create-organization endpoint.
type CreateOrganizationData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// isValidAccessToken checks if the provided access token is valid.
// Replace this function with your actual access token validation logic.
func isValidAccessToken(accessToken string) bool {
	// Your access token validation logic goes here
	// For demonstration purposes, always return true
	return true
}

// createOrganization creates a new organization and returns the organization ID.
// Replace this function with your actual organization creation logic.
func createOrganization(name, description string) string {
	// Your organization creation logic goes here
	// For demonstration purposes, generating a sample organization ID
	return "sample_organization_id"
}