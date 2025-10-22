// Package routes provides the HTTP route definitions for the AIpply.
package routes

import (
	"context"

	"github.com/DeleMike/AIpply/api/metrics"
	"github.com/DeleMike/AIpply/api/service"
	"github.com/gin-gonic/gin"
)

// GenerateCV handles the HTTP request to generate a CV based on
// user-provided answers and job details. It processes the input,
// interacts with the AI service, and returns a structured CV
// as part of the JSON response.
func GenerateCV(c *gin.Context) {
	handleLLMRequest(c, &GenerateCVRequest{}, func(ctx context.Context, request *GenerateCVRequest) (any, error) {

		cv, err := service.ProcessUserPrepAnswers(c, service.LLMClient, request.JobDescription, request.Answers)

		if err != nil {
			return nil, err
		}

		metrics.IncrementCV()

		return CVResponsePayload{
			Message: "CV generated successfully!",
			CV:      cv,
		}, nil
	})
}
