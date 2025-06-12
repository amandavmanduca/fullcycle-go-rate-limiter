package rate_limiter

import (
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/interfaces"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/structs"
)

type limits struct {
	apiKeyRateLimit int
	ipRateLimit     int
}

type rateLimiterService struct {
	rateLimiterRepository interfaces.RateLimiterRepository
	limits                limits
}

func NewRateLimiterService(configs structs.Configs, repository interfaces.RateLimiterRepository) interfaces.RateLimiterService {
	return &rateLimiterService{
		rateLimiterRepository: repository,
		limits: limits{
			apiKeyRateLimit: configs.RateLimitApiKey,
			ipRateLimit:     configs.RateLimitInterval,
		},
	}
}
