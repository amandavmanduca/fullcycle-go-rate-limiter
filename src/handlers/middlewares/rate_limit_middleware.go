package middlewares

import (
	"fmt"
	"net/http"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/interfaces"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	RateLimiterService interfaces.RateLimiterService
}

func (m middleware) RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		apiKey := c.Request().Header.Get("API_KEY")
		fmt.Println(apiKey)
		fmt.Println(ip)
		if apiKey == "" {
			return c.JSON(http.StatusTooManyRequests, map[string]string{"message": "you have reached the maximum number of requests or actions allowed within a certain time frame"})
		}
		return next(c)
	}
}
