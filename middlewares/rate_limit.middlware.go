package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/responses"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	smemory "github.com/ulule/limiter/v3/drivers/store/memory"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimiter struct {
	handler gin.HandlerFunc
}

// @formattedRate = e.g; 5-H -> 5 req/hour,
// H-> Hour, M -> Minute, S -> Second
func NewRateLimiter(formattedRate string) *RateLimiter {
	var store limiter.Store
	if config.RedisClient != nil {

		redisStore, err := sredis.NewStoreWithOptions(config.RedisClient, limiter.StoreOptions{
			Prefix:   "limiter_jquest",
			MaxRetry: 3,
		})

		if err != nil {
			log.Printf("failed to create Redis limiter store: %v — falling back to memory", err)
		} else {
			store = redisStore
			log.Println("using Redis rate limiter store")
		}
	}

	// fallback
	if store == nil {
		store = smemory.NewStore()
		log.Println("using in-memory rate limiter store")
	}

	rate, err := limiter.NewRateFromFormatted(formattedRate)
	if err != nil {
		log.Printf("invalid rate format %q: %v — defaulting to 4-H", formattedRate, err)
		rate, _ = limiter.NewRateFromFormatted("50-M")
	}

	instance := limiter.New(store, rate)

	mw := mgin.NewMiddleware(instance,
		mgin.WithLimitReachedHandler(func(c *gin.Context) {
			res := responses.Failure[any](http.StatusTooManyRequests, "Too many requests.")
			c.JSON(http.StatusTooManyRequests, res)
		}),
	)

	return &RateLimiter{handler: mw}
}

func (r *RateLimiter) Middleware() gin.HandlerFunc {
	return r.handler
}
