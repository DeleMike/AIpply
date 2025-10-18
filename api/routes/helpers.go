// Package routes provides the HTTP route definitions for the AIpply.
package routes

import (
	"context"
	"net/http"

	"github.com/DeleMike/AIpply/api/apierrors"
	"github.com/DeleMike/AIpply/api/service"
	"github.com/gin-gonic/gin"
)

// handleLLMRequest centralizes JSON binding, LLM readiness check,
// and standardized error handling for LLM-based routes.
// It accepts a generic payload type and a handler function to process the request.
func handleLLMRequest[T any](c *gin.Context, payload *T, handler func(context.Context, *T) (any, error)) {
	if err := c.ShouldBindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apierrors.ErrInvalidRequest})
		return
	}

	if service.LLMClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apierrors.ErrLLMNotInitialized})
		return
	}

	result, err := handler(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apierrors.ErrWeCouldNotProcessRequest})
		return
	}

	c.JSON(http.StatusOK, result)
}
