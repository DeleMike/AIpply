// Package api provides API routing and middleware setup for AIpply.
package api

import (
	"fmt"
	"log"

	"github.com/DeleMike/AIpply/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// StartUpServer initializes and starts the Gin server.
func StartUpServer() {
	// Set Gin mode based on environment.
	switch viper.GetString("server.gin_mode") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode) // fallback
	}

	fmt.Println("Server Gin Mode:", viper.GetString("server.gin_mode"))

	r := gin.Default()

	// Add custom middleware.
	r.Use(middleware.SetupCORS())

	// Setup API routes.
	SetupRouter(r)

	// Start the server.
	startServer(r)
}

// startServer starts the HTTP server.
func startServer(r *gin.Engine) {
	port := viper.GetInt("server.port")
	hostPort := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on %s...", hostPort)
	log.Fatal(r.Run(hostPort))
}
