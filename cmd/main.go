// cmd/main.go

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mina-droid/Ideanest-Software-Engineering-Task/pkg/api/middleware"
	"github.com/mina-droid/Ideanest-Software-Engineering-Task/pkg/api/routes"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Apply middleware globally
	router.Use(middleware.AuthMiddleware())

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
