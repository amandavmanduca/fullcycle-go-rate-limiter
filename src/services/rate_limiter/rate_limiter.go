package rate_limiter

import (
	"time"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/interfaces"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/structs"
)

type limits struct {
	apiKeyRateLimit int
	apiKeyDuration  time.Duration
	ipRateLimit     int
	ipDuration      time.Duration
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
			apiKeyDuration:  configs.RateLimitApiKeyInterval,
			ipRateLimit:     configs.RateLimitIp,
			ipDuration:      configs.RateLimitIpInterval,
		},
	}
}
