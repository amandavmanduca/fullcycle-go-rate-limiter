package interfaces

import (
	"context"
	"time"
)

type RateLimiterService interface {
	CheckRateLimit(ctx context.Context, ip string, apiKey string) error
}

type RateLimiterRepository interface {
	Increment(ctx context.Context, key string, expiration time.Duration) (int64, error)
}

type Redis interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Incr(ctx context.Context, key string, expiration time.Duration) (int64, error)
}
