// Package service provides core business logic and operations for the AIpply application.
package service

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"google.golang.org/genai"
)

// LLMClient is a shared Gemini API client instance used throughout the application.
var LLMClient *genai.Client

// float32Ptr returns a pointer to the given float32 value.
func float32Ptr(f float32) *float32 {
	return &f
}

// InitLLMService initializes the Gemini LLM client with the provided API key.
// It stores the client in a global variable for later reuse.
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

// ProcessUserPayload takes a job description and experience level,
// sends them to the Gemini model, and returns a list of generated interview questions.
func ProcessUserPayload(ctx context.Context, client *genai.Client, jobDescription string, experienceLevel string) ([]string, error) {
	model := "gemini-2.0-flash"

	prompt := fmt.Sprintf(
		Prompt,
		experienceLevel,
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

// parseQuestions cleans and structures raw model output into a slice of questions.
func parseQuestions(output string) []string {
	var result []string
	lines := strings.Split(output, "\n")

	// If the first line isn't a question, skip it
	if len(lines) > 0 && !strings.Contains(lines[0], "?") {
		lines = lines[1:]
	}

	// This regex finds lines that may start with numbers, bullets, or asterisks.
	leadingJunk := regexp.MustCompile(`^\s*([\d\.\-\*•\)]+\s*)*(.*)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			matches := leadingJunk.FindStringSubmatch(line)
			if len(matches) > 2 && matches[2] != "" {
				result = append(result, matches[2])
			}
		}
	}

	return result
}
