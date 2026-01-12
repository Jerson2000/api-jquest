package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cache/persistence"
)

var (
	CacheStore persistence.CacheStore
	onceCache  sync.Once
)

func configInitCacheStore() {
	onceCache.Do(func() {
		redisURL := os.Getenv("REDIS_URL")

		if redisURL != "" {
			store := persistence.NewRedisCacheWithURL(redisURL, time.Minute)
			if store != nil {
				CacheStore = store
				log.Println("INFO: using Redis cache store")
				return
			}
			log.Println("WARN: failed to init Redis cache — falling back to in-memory")
		}

		CacheStore = persistence.NewInMemoryStore(time.Second)
		log.Println("INFO: using in-memory cache store")
	})
}
