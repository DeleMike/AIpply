package routes

import (
	"net/http"

	"github.com/DeleMike/AIpply/api/metrics"
	"github.com/gin-gonic/gin"
)

// GetMetrics send merics about the current CVs and cover letters generated
func GetMetrics(c *gin.Context) {
	response := metrics.GetMetrics()
	c.JSON(http.StatusOK, response)

}
