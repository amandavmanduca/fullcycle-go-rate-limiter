package interfaces

import (
	"time"

	"github.com/labstack/echo/v4"
)

type RateLimiterService interface {
	RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type RateLimiterRepository interface{}

type Redis interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, expiration time.Duration) error
}
