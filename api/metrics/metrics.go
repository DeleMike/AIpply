// Package metrics provides Prometheus-style metrics tracking and exposure.
package metrics

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
)

const (
	keyCVGenerated          = "cv_generated"
	keyCoverLetterGenerated = "cover_letter_generated"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

// InitRedis now accepts the connection parameters
func InitRedis(addr string, password string, db int) {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		if err := client.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("Could not connect to Redis: %v", err)
		}
		log.Println("Successfully connected to Redis")
		redisClient = client
	})
}

// IncrementCV atomically adds one to the CV counter
func IncrementCV() {
	if redisClient == nil {
		log.Println("Redis client not initialized, skipping metric increment")
		return
	}

	err := redisClient.Incr(context.Background(), keyCVGenerated).Err()
	if err != nil {
		log.Printf("Failed to increment %s: %v", keyCVGenerated, err)
	}
}

// IncrementCoverLetter atomically adds one to the Cover Letter counter
func IncrementCoverLetter() {
	if redisClient == nil {
		log.Println("Redis client not initialized, skipping metric increment")
		return
	}

	err := redisClient.Incr(context.Background(), keyCoverLetterGenerated).Err()
	if err != nil {
		log.Printf("Failed to increment %s: %v", keyCoverLetterGenerated, err)
	}
}

// MetricResponse defines the JSON structure for our /metrics endpoint
type MetricResponse struct {
	CVGenerated          int64 `json:"cv_generated"`
	CoverLetterGenerated int64 `json:"cover_letter_generated"`
}

// GetMetrics safely reads the current CV and Cover Letter counts
func GetMetrics() MetricResponse {
	if redisClient == nil {
		log.Println("Redis client not initialized, returning zero metrics")
		return MetricResponse{}
	}

	getMetricValue := func(key string) int64 {
		val, err := redisClient.Get(context.Background(), key).Result()

		if err == redis.Nil {
			return 0
		}

		if err != nil {
			log.Printf("Failed to get metric %s: %v", key, err)
			return 0
		}

		count, _ := strconv.ParseInt(val, 10, 64)
		return count
	}

	cvCount := getMetricValue(keyCVGenerated)
	clCount := getMetricValue(keyCoverLetterGenerated)

	return MetricResponse{
		CVGenerated:          cvCount,
		CoverLetterGenerated: clCount,
	}
}
