// pkg/api/routes/routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mina-droid/Ideanest-Software-Engineering-Task/pkg/api/handlers"
)

// SetupRoutes initializes the API routes.
func SetupRoutes(router *gin.Engine) {
	// Sign up endpoint
	router.POST("/signup", handlers.SignUpHandler)

	// Sign in endpoint
	router.POST("/signin", handlers.SignInHandler)

	// Refresh token endpoint
	router.POST("/refresh-token", handlers.RefreshTokenHandler)

	// Create organization endpoint
	router.POST("/organization", handlers.CreateOrganizationHandler)

	// Read organization endpoint
	router.GET("/organization/:organization_id", handlers.ReadOrganizationHandler)

	// Read all organizations endpoint
	router.GET("/organization", handlers.ReadAllOrganizationsHandler)

	// Update organization endpoint
	router.PUT("/organization/:organization_id", handlers.UpdateOrganizationHandler)

	// Delete organization endpoint
	router.DELETE("/organization/:organization_id", handlers.DeleteOrganizationHandler)

	// Invite user to organization endpoint
	router.POST("/organization/:organization_id/invite", handlers.InviteUserToOrganizationHandler)
}
