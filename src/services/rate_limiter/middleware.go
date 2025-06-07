package rate_limiter

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s rateLimiterService) RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
