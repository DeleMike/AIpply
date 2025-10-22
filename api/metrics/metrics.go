package metrics

import (
	"sync"
)

// Holds our application's metrics
type store struct {
	mu                   sync.RWMutex
	cvGenerated          int64
	coverLetterGenerated int64
}

// Global instance of our metric store
var metricStore = &store{}

// IncrementCV atomically adds one to the CV counter
func IncrementCV() {
	metricStore.mu.Lock()
	metricStore.cvGenerated++
	metricStore.mu.Unlock()
}

// IncrementCoverLetter atomically adds one to the Cover Letter counter
func IncrementCoverLetter() {
	metricStore.mu.Lock()
	metricStore.coverLetterGenerated++
	metricStore.mu.Unlock()
}

// MetricsResponse defines the JSON structure for our /metrics endpoint
type MetricResponse struct {
	CVGenerated          int64 `json:"cv_generated"`
	CoverLetterGenerated int64 `json:"cover_letter_generated"`
}

// GetMetrics safely reads the current counts
func GetMetrics() MetricResponse {
	metricStore.mu.RLock()
	defer metricStore.mu.RUnlock()
	return MetricResponse{
		CVGenerated:          metricStore.cvGenerated,
		CoverLetterGenerated: metricStore.coverLetterGenerated,
	}
}
