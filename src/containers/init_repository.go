package containers

import (
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers/repository"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/interfaces"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/repositories/rate_limiter"
)

func NewRepositoryContainer(redis interfaces.Redis) repository.RepositoryContainer {
	return repository.RepositoryContainer{
		RateLimiterRepository: rate_limiter.NewRateLimiterRepository(redis),
	}
}
