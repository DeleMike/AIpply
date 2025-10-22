// Package routes provides the HTTP route definitions for the AIpply.
package routes

import (
	"context"

	"github.com/DeleMike/AIpply/api/metrics"
	"github.com/DeleMike/AIpply/api/service"
	"github.com/gin-gonic/gin"
)

// GenerateQuestions handles the incoming job description payload,
// sends it to the LLM service to generate interview questions,
// and returns the questions in a structured JSON response.
func GenerateQuestions(c *gin.Context) {
	handleLLMRequest(c, &GenerateQuestionsRequest{}, func(ctx context.Context, request *GenerateQuestionsRequest) (any, error) {
		questions, err := service.ProcessUserPayload(c.Request.Context(), service.LLMClient, request.JobDescription, request.ExperienceLevel)

		if err != nil {
			return nil, err
		}

		metrics.IncrementCoverLetter()

		return QuestionsResponsePayload{
			Message:   "Questions generated successfully!",
			Questions: questions,
		}, nil

	})

}
