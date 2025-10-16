// Package routes provides the HTTP route definitions for the AIpply.
package routes

import (
	"net/http"

	"github.com/DeleMike/AIpply/api/service"
	"github.com/DeleMike/AIpply/api/utils"
	"github.com/gin-gonic/gin"
)

func GenerateQuestions(c *gin.Context) {
	// get payload of job
	var request GenerateQuestionsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidRequest})
		return
	}

	client := service.LLMClient
	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrLLMNotInitialized})
		return
	}

	// use provided data and send to the AI
	questions, err := service.ProcessUserPayload(c, client, request.JobDescription, request.ExperienceLevel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrWeCouldNotProcessRequest})
	}

	response := QuestionsResponsePayload{
		Message:   "Questions generated successfully!",
		Questions: questions,
	}

	c.JSON(http.StatusOK, response)

}
