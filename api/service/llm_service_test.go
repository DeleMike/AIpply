package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestParseQuestions tests the question parsing logic with various inputs
func TestParseQuestions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "numbered questions with dots",
			input: "1. What is your experience with Go?\n2. How do you handle errors?\n3. Explain interfaces?",
			expected: []string{
				"What is your experience with Go?",
				"How do you handle errors?",
				"Explain interfaces?",
			},
		},
		{
			name:  "questions with asterisks",
			input: "* What is your experience?\n* How do you test?",
			expected: []string{
				"What is your experience?",
				"How do you test?",
			},
		},
		{
			name:  "questions with bullets",
			input: "• First question?\n• Second question?",
			expected: []string{
				"First question?",
				"Second question?",
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "whitespace only",
			input:    "   \n\n   \n",
			expected: []string{},
		},
		{
			name:  "mixed formatting",
			input: "1) What is Go?\n2) Explain testing?\n- What about mocking?",
			expected: []string{
				"What is Go?",
				"Explain testing?",
				"What about mocking?",
			},
		},
		{
			name:  "with header line (should skip)",
			input: "Here are your questions:\n1. What is your name?\n2. Where do you work?",
			expected: []string{
				"What is your name?",
				"Where do you work?",
			},
		},
		{
			name:  "questions without numbering",
			input: "What is your experience?\nHow do you handle errors?",
			expected: []string{
				"What is your experience?",
				"How do you handle errors?",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseQuestions(tt.input)

			assert.Equal(t, len(tt.expected), len(result),
				"Expected %d questions but got %d", len(tt.expected), len(result))

			for i, expected := range tt.expected {
				if i < len(result) {
					assert.Equal(t, expected, result[i],
						"Question %d mismatch", i)
				}
			}
		})
	}
}

// TestFloat32Ptr tests the helper function
func TestFloat32Ptr(t *testing.T) {
	tests := []struct {
		name  string
		input float32
	}{
		{"zero", 0.0},
		{"positive", 0.5},
		{"negative", -1.5},
		{"one", 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := float32Ptr(tt.input)

			require.NotNil(t, result, "Result should not be nil")
			assert.Equal(t, tt.input, *result, "Value should match input")
		})
	}
}

// Benchmark for parseQuestions to ensure performance
func BenchmarkParseQuestions(b *testing.B) {
	input := `Here are your interview questions:
1. What is your experience with Go?
2. How do you handle concurrent programming?
3. Explain goroutines and channels?
4. What testing frameworks have you used?
5. How do you ensure code quality?`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseQuestions(input)
	}
}
