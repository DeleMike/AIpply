// services provides core business logic and operations for the AIpply application.
package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/genai"
)

var LLMClient *genai.Client

func InitLLMService(ctx context.Context, apiKey string) error {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return fmt.Errorf("error initializing LLM client: %w", err)
	}
	LLMClient = client
	return nil
}

func float32Ptr(f float32) *float32 {
	return &f
}

// ProcessUserPayload takes the genai client and the job description,
// and asks Gemini to generate interview questions.
func ProcessUserPayload(ctx context.Context, client *genai.Client, jobDescription string, experienceLevel string) ([]string, error) {
	model := "gemini-2.0-flash"

	jobCtx := ExtractJobContext(jobDescription)

	prompt := fmt.Sprintf(
		Prompt,
		experienceLevel,
		jobCtx.Role,
		jobCtx.Company,
		strings.Join(jobCtx.Keywords, ", "),
		jobDescription,
	)

	resp, err := client.Models.GenerateContent(ctx, model,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			Temperature:     float32Ptr(0.6),
			TopP:            float32Ptr(0.9),
			MaxOutputTokens: 900,
		},
	)
	if err != nil {
		log.Printf("error generating content: %v", err)
		return nil, err
	}

	if len(resp.Candidates) == 0 {
		return nil, fmt.Errorf("no candidates returned from model")
	}

	// Extract the text output
	var output strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		output.WriteString(string(part.Text))
	}

	questions := parseQuestions(string(output.String()))

	return questions, nil
}

// parseQuestions converts the raw LLM text into a []string slice.
func parseQuestions(output string) []string {
	var result []string
	for line := range strings.SplitSeq(output, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			// Remove leading numbers, bullets, or extra dots/spaces
			line = strings.TrimLeft(line, "0123456789.-)• ")
			result = append(result, line)
		}
	}

	return result
}
