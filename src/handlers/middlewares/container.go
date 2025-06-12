package middlewares

import "github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers/service"

func NewMiddleware(s service.ServiceContainer) middleware {
	return middleware{
		RateLimiterService: s.RateLimiterService,
	}
}
