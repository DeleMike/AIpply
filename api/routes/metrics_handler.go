package routes

import (
	"net/http"

	"github.com/DeleMike/AIpply/api/metrics"
	"github.com/gin-gonic/gin"
)

func GetMetrics(c *gin.Context) {
	response := metrics.GetMetrics()
	c.JSON(http.StatusOK, response)

}
