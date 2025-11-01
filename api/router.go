package api

import (
	// "github.com/DeleMike/AIpply/api/middleware"
	"github.com/DeleMike/AIpply/api/routes"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all API routes with versioning.
func SetupRouter(r *gin.Engine) {
	r.GET("/", routes.Health)

	// API versioning
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// v1.POST("/generate-questions", middleware.RateLimit(), routes.GenerateQuestions)
			// v1.POST("/generate-cv", middleware.RateLimit(), routes.GenerateCV)
			// v1.POST("/generate-cover-letter", middleware.RateLimit(), routes.GenerateCoverLetter)
			v1.POST("/generate-questions", routes.GenerateQuestions)
			v1.POST("/generate-cv", routes.GenerateCV)
			v1.POST("/generate-cover-letter", routes.GenerateCoverLetter)
			v1.GET("/metrics", routes.GetMetrics)

		}
	}
}
