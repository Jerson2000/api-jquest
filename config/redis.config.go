package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func configRedisClient() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Println("WARN: REDIS_URL environment variable is not set — Redis disabled")
		return
	}

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Printf("WARN: failed to parse REDIS_URL: %v — Redis disabled\n", err)
		return
	}

	client := redis.NewClient(opts)

	// Try to ping Redis with a short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("WARN: could not connect to Redis: %v — Redis disabled\n", err)
		return
	}

	RedisClient = client
	log.Println("INFO: connected to Redis successfully")
}
