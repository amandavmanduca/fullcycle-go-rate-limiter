package rate_limiter

import "github.com/amandavmanduca/fullcycle-go-rate-limiter/src/interfaces"

type rateLimiterService struct {
	rateLimiterRepository interfaces.RateLimiterRepository
}

func NewRateLimiterService(repository interfaces.RateLimiterRepository) interfaces.RateLimiterService {
	return &rateLimiterService{
		rateLimiterRepository: repository,
	}
}
