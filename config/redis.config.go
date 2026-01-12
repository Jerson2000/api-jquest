package config

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	onceRedis   sync.Once
)

func configRedisClient() {
	onceRedis.Do(func() {
		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {
			log.Println("WARN: REDIS_URL not set — Redis disabled")
			return
		}

		opts, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Printf("WARN: invalid REDIS_URL: %v — Redis disabled\n", err)
			return
		}

		client := redis.NewClient(opts)

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := client.Ping(ctx).Err(); err != nil {
			_ = client.Close()
			log.Printf("WARN: Redis connection failed: %v — Redis disabled\n", err)
			return
		}

		RedisClient = client
		log.Println("INFO: Redis connected")
	})
}
