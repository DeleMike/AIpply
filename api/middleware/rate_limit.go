// Package middleware provides middleware setup for AIpply.
package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/DeleMike/AIpply/api/metrics"
	"github.com/gin-gonic/gin"
)

const (
	rateLimitPerDay = 18 // 6 full cycles per day (3 endpoints × 6 times)
	redisPrefix     = "rate_limit:"
)

// RateLimit is a Gin middleware to limit requests per IP.
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		redisClient := metrics.GetRedisClient()

		if redisClient == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Rate limit service unavailable"})
			return
		}

		ip := c.ClientIP()
		now := time.Now()
		key := redisPrefix + ip + ":" + now.Format("2006-01-02")
		ctx := context.Background()

		// get current trial count
		count, err := redisClient.Get(ctx, key).Int64()
		if err != nil && err.Error() != "redis: nil" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check rate limit"})
			return
		}

		// check limit
		if count >= rateLimitPerDay {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Daily limit exceeded",
				"message": "You've reached your daily limit of 6 generations. Please try again in 24 hours.",
			})
			return
		}

		// Compute time until midnight for expiry
		midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		expiry := time.Until(midnight)

		// update redis data
		pipe := redisClient.TxPipeline()
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, expiry)
		_, err = pipe.Exec(ctx)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rate limit"})
			return
		}

		c.Next()

	}
}
