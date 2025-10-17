// Package apierrors defines common error constants for the AIpply API.
package apierrors

const (
	// ErrInvalidRequest indicates that the request payload is invalid.
	ErrInvalidRequest = "Invalid request payload"
	// ErrLLMNotInitialized indicates the LLM service initialised well
	ErrLLMNotInitialized = "LLM service not initialized"
	// ErrWeCouldNotProcessRequest indicates we could not process request
	ErrWeCouldNotProcessRequest = "We could not process request"
)
