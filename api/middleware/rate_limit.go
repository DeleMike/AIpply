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
	rateLimitPerDay = 10
	expiry          = 24 * time.Hour
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
		key := redisPrefix + ip
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
				"error":   "You have exceeded your daily limit of 5 requests.",
				"message": "Please try again in 24 hours.",
			})
			return
		}

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
