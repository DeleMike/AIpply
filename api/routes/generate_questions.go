// Package routes provides the HTTP route definitions for the AIpply.
package routes

import (
	"net/http"

	"github.com/DeleMike/AIpply/api/apierrors"
	"github.com/DeleMike/AIpply/api/service"
	"github.com/gin-gonic/gin"
)

// GenerateQuestions handles the incoming job description payload,
// sends it to the LLM service to generate interview questions,
// and returns the questions in a structured JSON response.
func GenerateQuestions(c *gin.Context) {
	var request GenerateQuestionsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apierrors.ErrInvalidRequest})
		return
	}

	if service.LLMClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apierrors.ErrLLMNotInitialized})
		return
	}

	questions, err := service.ProcessUserPayload(c.Request.Context(), service.LLMClient, request.JobDescription, request.ExperienceLevel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apierrors.ErrWeCouldNotProcessRequest})
		return
	}

	response := QuestionsResponsePayload{
		Message:   "Questions generated successfully!",
		Questions: questions,
	}

	c.JSON(http.StatusOK, response)

}
