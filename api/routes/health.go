// Package routes provides the HTTP route definitions for the AIpply.
package routes

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the AIpply API!",
	})
}
