// Package service provides core business logic and operations for the AIpply application.
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/DeleMike/AIpply/api/models"
	"google.golang.org/genai"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// isRateLimitError checks if the error is a rate limit error (429 or RESOURCE_EXHAUSTED)
func isRateLimitError(err error) bool {
	if err == nil {
		return false
	}

	// Check for gRPC status code
	if st, ok := status.FromError(err); ok {
		return st.Code() == codes.ResourceExhausted
	}

	// Check for string patterns
	errStr := err.Error()
	return strings.Contains(errStr, "429") ||
		strings.Contains(errStr, "rate limit") ||
		strings.Contains(errStr, "quota")
}

// generateContentWithRetry attempts to generate content with automatic retry and fallback
func generateContentWithRetry(
	ctx context.Context,
	client *genai.Client,
	primaryModel string,
	fallbackModel string,
	prompt string,
	config *genai.GenerateContentConfig,
) (*genai.GenerateContentResponse, error) {
	maxRetries := 3
	baseDelay := 2 * time.Second

	// Try primary model with retries
	log.Printf("Using %s for processing...", primaryModel)
	for attempt := range maxRetries {
		resp, err := client.Models.GenerateContent(ctx, primaryModel,
			genai.Text(prompt),
			config,
		)

		if err == nil {
			return resp, nil
		}

		// If rate limited, wait and retry
		if isRateLimitError(err) {
			if attempt < maxRetries-1 {
				delay := baseDelay * time.Duration(attempt+1)
				log.Printf("Rate limit hit on %s, retrying in %v (attempt %d/%d)",
					primaryModel, delay, attempt+1, maxRetries)
				time.Sleep(delay)
				continue
			}
			log.Printf("Rate limit exhausted on %s after %d attempts, falling back to %s",
				primaryModel, maxRetries, fallbackModel)
			break
		}

		// For other errors, return immediately
		return nil, err
	}

	// Fallback to secondary model
	if fallbackModel != "" && fallbackModel != primaryModel {
		log.Printf("Using fallback model: %s", fallbackModel)
		resp, err := client.Models.GenerateContent(ctx, fallbackModel,
			genai.Text(prompt),
			config,
		)
		if err != nil {
			return nil, fmt.Errorf("fallback model %s also failed: %w", fallbackModel, err)
		}
		return resp, nil
	}

	return nil, fmt.Errorf("all attempts failed for model %s", primaryModel)
}

// ProcessUserPayload takes a job description and experience level,
// sends them to the Gemini model, and returns a list of generated interview questions.
func ProcessUserPayload(ctx context.Context, client *genai.Client, jobDescription string, experienceLevel string) ([]string, error) {
	model := "gemini-2.0-flash"

	prompt := fmt.Sprintf(
		QuestionPrompt,
		experienceLevel,
		jobDescription,
	)

	resp, err := client.Models.GenerateContent(ctx, model,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			Temperature:     float32Ptr(0.6),
			TopP:            float32Ptr(0.9),
			MaxOutputTokens: 1024,
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

// ProcessUserPrepAnswers is used to build the CV for the user based on their answered questions
func ProcessUserPrepAnswers(ctx context.Context, client *genai.Client, jobDescription string, answers []models.AnswerPair) (string, error) {
	primaryModel := "gemini-2.5-pro"
	fallbackModel := "gemini-2.5-flash"

	answersJSON, err := json.Marshal(answers)
	if err != nil {
		return "", fmt.Errorf("failed to marshal answers: %w", err)
	}

	prompt := fmt.Sprintf(
		CVPrompt,
		jobDescription,
		string(answersJSON),
	)

	resp, err := generateContentWithRetry(ctx, client, primaryModel, fallbackModel, prompt,
		&genai.GenerateContentConfig{
			Temperature:     float32Ptr(0.2),
			TopP:            float32Ptr(0.9),
			MaxOutputTokens: 10000,
		},
	)

	if err != nil {
		log.Printf("error generating content from LLM: %v", err)
		return "", fmt.Errorf("error from LLM: %w", err)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no candidates returned from model")
	}

	var cvText strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		cvText.WriteString(string(part.Text))
	}

	if cvText.Len() == 0 {
		return "", fmt.Errorf("empty response from model")
	}

	return cvText.String(), nil
}

// ProcessUserAnswersForCoverLetter is used to generate the cover letter of the user
func ProcessUserAnswersForCoverLetter(ctx context.Context, client *genai.Client, jobDescription string, answers []models.AnswerPair) (string, error) {
	primaryModel := "gemini-2.5-pro"
	fallbackModel := "gemini-2.5-flash"

	answersJSON, err := json.Marshal(answers)
	if err != nil {
		return "", fmt.Errorf("failed to marshal answers: %w", err)
	}

	prompt := fmt.Sprintf(
		CoverLetterPrompt,
		jobDescription,
		string(answersJSON),
	)

	resp, err := generateContentWithRetry(ctx, client, primaryModel, fallbackModel, prompt,
		&genai.GenerateContentConfig{
			Temperature:     float32Ptr(0.2),
			TopP:            float32Ptr(0.9),
			MaxOutputTokens: 10000,
		},
	)

	if err != nil {
		log.Printf("error generating content from LLM: %v", err)
		return "", fmt.Errorf("error from LLM: %w", err)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no candidates returned from model")
	}

	var cvText strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		cvText.WriteString(string(part.Text))
	}

	if cvText.Len() == 0 {
		return "", fmt.Errorf("empty response from model")
	}

	return cvText.String(), nil
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
