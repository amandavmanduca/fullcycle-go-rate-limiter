package middlewares

import (
	"net/http"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/interfaces"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	RateLimiterService interfaces.RateLimiterService
	ValidApiKey        string
}

func (m middleware) RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/favicon.ico" {
			return next(c)
		}
		ctx := c.Request().Context()
		ip := c.RealIP()
		foundKey := c.Request().Header.Get("API_KEY")
		err := m.RateLimiterService.CheckRateLimit(ctx, ip, validApiKey(foundKey, m.ValidApiKey))
		if err != nil {
			return c.JSON(http.StatusTooManyRequests, map[string]string{"message": err.Error()})
		}
		return next(c)
	}
}

func validApiKey(foundKey, apiKey string) string {
	if foundKey == "" {
		return ""
	}
	if foundKey == apiKey {
		return foundKey
	}
	return ""
}
