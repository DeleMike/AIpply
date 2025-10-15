package api

import (
	"github.com/DeleMike/AIpply/api/routes"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all API routes with versioning.
func SetupRouter(r *gin.Engine) {
	r.GET("/", routes.Health)

	// API versioning
	// api := r.Group("/api")
	// {
	// 	v1 := api.Group("/v1")
	// 	{
	//
	// 	}
	// }
}
