package rate_limiter

import (
	"context"
	"sync"
	"testing"
	"time"

	// "github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/redis"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/repositories/rate_limiter"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/structs"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCheckRateLimit(t *testing.T) {
	configs := structs.Configs{
		ApiKey:                  "secretkey",
		RateLimitIp:             5,
		RateLimitApiKey:         10,
		RateLimitIpInterval:     time.Second * 3,
		RateLimitApiKeyInterval: time.Second * 3,
	}
	t.Run("should accept all ip requests when rate limit is not exceeded", func(t *testing.T) {
		testutils.WithDB(t, "should accept all ip requests when rate limit is not exceeded", func(t *testing.T, redis *redis.Redis) {
			repo := rate_limiter.NewRateLimiterRepository(redis)
			service := NewRateLimiterService(configs, repo)
			ctx := context.Background()
			ip := "127.0.0.2"
			apiKey := ""
			wg := sync.WaitGroup{}

			size := 5
			wg.Add(size)
			results := make(chan error, size)

			for i := 0; i < size; i++ {
				go func() {
					defer wg.Done()
					err := service.CheckRateLimit(ctx, ip, apiKey)
					if err != nil {
						results <- err
					}
				}()
			}

			wg.Wait()
			close(results)

			assert.Equal(t, 0, len(results))
		})
	})
	t.Run("should return error when ip rate limit is exceeded", func(t *testing.T) {
		testutils.WithDB(t, "should return error when ip rate limit is exceeded", func(t *testing.T, redis *redis.Redis) {
			repo := rate_limiter.NewRateLimiterRepository(redis)
			service := NewRateLimiterService(configs, repo)
			ctx := context.Background()
			ip := "127.0.0.1"
			apiKey := ""
			wg := sync.WaitGroup{}

			size := 10
			wg.Add(size)
			results := make(chan error, size)

			for i := 0; i < size; i++ {
				go func() {
					defer wg.Done()
					err := service.CheckRateLimit(ctx, ip, apiKey)
					if err != nil {
						results <- err
					}
				}()
			}

			wg.Wait()
			close(results)

			assert.Equal(t, 5, len(results))
		})
	})
	t.Run("should accept all api key requests when rate limit is not exceeded", func(t *testing.T) {
		testutils.WithDB(t, "should accept all api key requests when rate limit is not exceeded", func(t *testing.T, redis *redis.Redis) {
			repo := rate_limiter.NewRateLimiterRepository(redis)
			service := NewRateLimiterService(configs, repo)
			ctx := context.Background()
			ip := "127.0.0.3"
			apiKey := "secretkey"
			wg := sync.WaitGroup{}

			size := 10
			wg.Add(size)
			results := make(chan error, size)

			for i := 0; i < size; i++ {
				go func() {
					defer wg.Done()
					err := service.CheckRateLimit(ctx, ip, apiKey)
					if err != nil {
						results <- err
					}
				}()
			}

			wg.Wait()
			close(results)

			assert.Equal(t, 0, len(results))
		})
	})
	t.Run("should return error when api key rate limit is exceeded", func(t *testing.T) {
		testutils.WithDB(t, "should return error when api key rate limit is exceeded", func(t *testing.T, redis *redis.Redis) {
			repo := rate_limiter.NewRateLimiterRepository(redis)
			key := "secretkey2"
			configs.ApiKey = key
			service := NewRateLimiterService(configs, repo)
			ctx := context.Background()
			ip := "127.0.0.4"
			apiKey := key
			wg := sync.WaitGroup{}

			size := 15
			wg.Add(size)
			results := make(chan error, size)

			for i := 0; i < size; i++ {
				go func() {
					defer wg.Done()
					err := service.CheckRateLimit(ctx, ip, apiKey)
					if err != nil {
						results <- err
					}
				}()
			}

			wg.Wait()
			close(results)

			assert.Equal(t, 5, len(results))
		})
	})
}
