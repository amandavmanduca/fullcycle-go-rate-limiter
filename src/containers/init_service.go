package containers

import (
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers/repository"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers/service"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/services/rate_limiter"
)

func NewServiceContainer(repositories repository.RepositoryContainer) service.ServiceContainer {
	return service.ServiceContainer{
		RateLimiterService: rate_limiter.NewRateLimiterService(repositories.RateLimiterRepository),
	}
}
