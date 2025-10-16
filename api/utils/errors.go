// Package utils contains utility functions and error handling mechanisms for the AIpply application.
package utils

const (
	// ErrInvalidRequest indicates that the request payload is invalid.
	ErrInvalidRequest           = "Invalid request payload"
	ErrLLMNotInitialized        = "LLM service not initialized"
	ErrWeCouldNotProcessRequest = "We could not process request"
)
