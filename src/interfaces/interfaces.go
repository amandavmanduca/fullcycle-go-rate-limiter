package interfaces

import (
	"time"
)

type RateLimiterService interface {
}

type RateLimiterRepository interface{}

type Redis interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, expiration time.Duration) error
}
