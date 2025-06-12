package interfaces

import (
	"context"
	"time"
)

type RateLimiterService interface {
}

type RateLimiterRepository interface{}

type Redis interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}
