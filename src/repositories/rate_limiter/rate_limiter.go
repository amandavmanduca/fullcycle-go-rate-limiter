package rate_limiter

import "github.com/amandavmanduca/fullcycle-go-rate-limiter/src/interfaces"

type rateLimiterRepository struct {
	redis interfaces.Redis
}

func NewRateLimiterRepository(redis interfaces.Redis) interfaces.RateLimiterRepository {
	return &rateLimiterRepository{
		redis: redis,
	}
}
