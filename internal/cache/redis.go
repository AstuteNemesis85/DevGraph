package cache

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	// For Render deployment - use REDIS_URL
	redisURL := os.Getenv("REDIS_URL")
	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatal("Failed to parse Redis URL:", err)
		}
		return redis.NewClient(opt)
	}

	// For local development - use REDIS_ADDR
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379" // default
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password for local dev
		DB:       0,
	})
}
