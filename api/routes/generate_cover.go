// Package routes provides the HTTP route definitions for the AIpply.
package routes

import (
	"context"

	"github.com/DeleMike/AIpply/api/service"
	"github.com/gin-gonic/gin"
)

// GenerateCoverLetter handles the HTTP request to generate a Cover Letter based on
// user-provided answers and job details. It processes the input,
// interacts with the AI service, and returns a great cover letter
// as part of the JSON response.
func GenerateCoverLetter(c *gin.Context) {
	handleLLMRequest(c, &GenerateCVRequest{}, func(ctx context.Context, request *GenerateCVRequest) (any, error) {

		coverLetter, err := service.ProcessUserAnswersForCoverLetter(c, service.LLMClient, request.JobDescription, request.Answers)

		if err != nil {
			return nil, err
		}

		return CVResponsePayload{
			Message: "Cover letter generated successfully!",
			CV:      coverLetter,
		}, nil
	})
}
