package middlewares

import (
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers/service"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/structs"
)

func NewMiddleware(configs structs.Configs, s service.ServiceContainer) middleware {
	return middleware{
		RateLimiterService: s.RateLimiterService,
		ValidApiKey:        configs.ApiKey,
	}
}
